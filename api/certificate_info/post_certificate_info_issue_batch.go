package certificateinfo

import (
	"encoding/json"
	"log"

	apiworker "cybercert-blockchain-api/apiworker"
	workermodels "cybercert-blockchain-api/apiworker/models"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func PostCertificateInfoIssueBatchImpl(params certinfoops.PostCertificateInfoIssueBatchParams) middleware.Responder {
	// Generate templateKey as template reference on blockchain
	// templateKey in `certificate_template` is equal to templateRef in `certificate_info`
	templateKey := uuid.New().String()
	log.Println("Generated templateKey:", templateKey)

	certParams := workermodels.CertificateBatchParam{
		OrganizationName: params.CertificateBatch.OrganizationName,
		CertTemplate: workermodels.CertificateTemplate{
			TemplateKey:    templateKey,
			TemplateSource: params.CertificateBatch.Template.TemplateSource,
			SourceType:     params.CertificateBatch.Template.SourceType,
			Version:        params.CertificateBatch.Template.Version,
			IssuerID:       params.CertificateBatch.Template.IssuerID,
			IssuerName:     params.CertificateBatch.Template.IssuerName,
		},
		CertDetails: []*workermodels.CertificateDetails{},
		CallbackUrl: params.CertificateBatch.CallbackURL,
	}

	for _, c := range params.CertificateBatch.CertificateDetails {
		certParams.CertDetails = append(certParams.CertDetails, &workermodels.CertificateDetails{
			CertificateID:        c.CertificateID,
			CertificateHolder:    c.CertificateHolder,
			TemplateRef:          templateKey,
			CourseName:           c.CourseName,
			ModuleName:           c.ModuleName,
			Email:                c.Email,
			IssuedAt:             c.IssuedAt,
			IssuerName:           c.IssuerName,
			CertificateSignature: c.CertificateSignature,
			Extras:               c.Extras,
		})
	}

	message, err := json.Marshal(certParams)
	if err != nil {
		return &certinfoops.PostCertificateInfoIssueBatchInternalServerError{
			Payload: &apimodels.APIResponse{Message: err.Error()},
		}
	}

	messageId := ""
	routingKey := viper.GetString("rabbitmq.routing_key")
	apiworker.EmitMiningJob(messageId, routingKey, message)

	return &certinfoops.PostCertificateInfoIssueBatchOK{
		Payload: &apimodels.APIResponse{
			Message: "OK",
		},
	}
}

var PostCertificateInfoIssueBatchHandler = certinfoops.PostCertificateInfoIssueBatchHandlerFunc(PostCertificateInfoIssueBatchImpl)
