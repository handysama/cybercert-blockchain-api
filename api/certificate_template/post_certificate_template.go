package certificatetemplate

import (
	"log"

	cert_template_api "cybercert-blockchain-api/chaincodeapi/certificate_template"
	apimodels "cybercert-blockchain-api/models"
	templateops "cybercert-blockchain-api/restapi/operations/certificate_template"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateTemplateImpl(params templateops.PostCertificateTemplateParams) middleware.Responder {
	result, err := cert_template_api.PutTemplate(
		params.CertificateTemplate.OrganizationName,
		params.CertificateTemplate.Template.TemplateKey,
		[]byte(params.CertificateTemplate.Template.TemplateSource),
		params.CertificateTemplate.Template.SourceType,
		params.CertificateTemplate.Template.Version,
		params.CertificateTemplate.Template.IssuerID,
		params.CertificateTemplate.Template.IssuerName,
	)

	if err != nil {
		log.Println("Error in PutTemplate:", err.Error())
		return &templateops.PostCertificateTemplateInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	if result.TransactionID == "" {
		log.Println("Error in PutTemplate: returned TransactionID is empty")
		return &templateops.PostCertificateTemplateInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: "Fail to put template on blockchain",
			},
		}
	}

	return &templateops.PostCertificateTemplateCreated{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostCertificateTemplateHandler = templateops.PostCertificateTemplateHandlerFunc(PostCertificateTemplateImpl)
