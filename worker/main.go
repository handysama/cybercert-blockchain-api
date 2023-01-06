package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	workermodels "cybercert-blockchain-api/apiworker/models"
)

func initConfig() {
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s\n", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func getDialUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s%s",
		viper.GetString("rabbitmq.scheme"),
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetString("rabbitmq.port"),
		viper.GetString("rabbitmq.vhost"),
	)
}

func getIssueCertificateUrl() string {
	api_host := viper.GetString("api_host")
	return fmt.Sprintf("%s/api/certificate/info/issue", api_host)
}

func getPutTemplateUrl() string {
	api_host := viper.GetString("api_host")
	return fmt.Sprintf("%s/api/certificate/template", api_host)
}

func Post(endpoint string, data []byte) ([]byte, error) {
	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", endpoint, body)
	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return bodyBytes, err
}

func invokeApi(apiUrl string, data []byte) (string, error) {
	response, err := Post(apiUrl, data)
	if err != nil {
		return "", err
	}

	parsedResponse := map[string]string{}
	err = json.Unmarshal(response, &parsedResponse)
	if err != nil {
		return "", err
	}

	return parsedResponse["transaction_id"], nil
}

func listenCertificateIssueBatch() {
	dialUrl := getDialUrl()
	conn, err := amqp.Dial(dialUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := viper.GetString("rabbitmq.exchange_issue_batch")
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	queueName := fmt.Sprintf("%s-%s", exchangeName, uuid.New().String())
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	routingKey := viper.GetString("rabbitmq.routing_key")

	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", routingKey)
	err = ch.QueueBind(
		q.Name,       // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no wait
		nil,          // args
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received task certificate issue batch. Message ID: %v, Length: %d", d.MessageId, len(d.Body))

			parsedBody := &workermodels.CertificateBatchParam{}
			err := json.Unmarshal(d.Body, parsedBody)
			if err != nil {
				log.Printf("Error in unmarshall query result: %s", err.Error())
				continue
			}

			templatePayload := map[string]interface{}{
				"organization_name": parsedBody.OrganizationName,
				"template":          parsedBody.CertTemplate,
			}

			templateBytes, err := json.Marshal(templatePayload)
			if err != nil {
				log.Printf("Error in marshal Template: %s", err.Error())
				continue
			}

			templateTransactionId, err := invokeApi(getPutTemplateUrl(), templateBytes)
			log.Println("Put template:", parsedBody.CertTemplate.TemplateKey, "Txn:", templateTransactionId, "Err:", err)

			if templateTransactionId == "" {
				log.Println("Error in Put Template: returned TransactionId is empty")
				continue
			}

			for _, c := range parsedBody.CertDetails {
				issuePayload := map[string]interface{}{
					"organization_name":   parsedBody.OrganizationName,
					"certificate_details": c,
				}

				payloadBytes, err := json.Marshal(issuePayload)
				if err != nil {
					log.Printf("Error in marshal CertDetails: %s", err.Error())
					continue
				}

				// Rate limit issuing certificate to avoid SPAM marked sender
				time.Sleep(2 * time.Second)

				issuerCertTransactionId, err := invokeApi(getIssueCertificateUrl(), payloadBytes)
				log.Println("Issuer Certificate:", c.CertificateID, "Txn:", issuerCertTransactionId, "Err:", err)

				// Notify platform-api for issued certificate
				notifyPayload := map[string]interface{}{
					"certificate_id":  c.CertificateID,
					"certificate_ref": issuerCertTransactionId,
					"template_ref":    parsedBody.CertTemplate.TemplateKey,
					"status":          "ISSUED",
				}
				notifyPayloadBytes, err := json.Marshal(notifyPayload)
				if err != nil {
					log.Printf("Error in marshal notify payload: %s", err.Error())
					continue
				}

				_, err = Post(parsedBody.CallbackUrl, notifyPayloadBytes)
				if err != nil {
					log.Printf("Failed to notify platform. certificate_id: %s, callback: %s", c.CertificateID, parsedBody.CallbackUrl)
				}
			}
		}
	}()

	log.Printf("Worker Certificate Issuer Queue started")
	<-forever
}

func main() {
	initConfig()
	listenCertificateIssueBatch()
}
