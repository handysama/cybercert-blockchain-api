package apiworker

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

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

func EmitMiningJob(jobId, routingKey string, message []byte) {
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

	err = ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			MessageId:   jobId,
			ContentType: "application/json",
			Body:        message,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent certificate issue batch request. Routing Key: %s, Message ID: %s, Length: %d",
		routingKey, jobId, len(message))
}
