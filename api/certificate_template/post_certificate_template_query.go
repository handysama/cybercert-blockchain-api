package certificatetemplate

import (
	"log"

	cert_template_api "cybercert-blockchain-api/chaincodeapi/certificate_template"
	apimodels "cybercert-blockchain-api/models"
	templateops "cybercert-blockchain-api/restapi/operations/certificate_template"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateTemplateQueryImpl(params templateops.PostCertificateTemplateQueryParams) middleware.Responder {
	result, err := cert_template_api.QueryTemplate(
		params.QueryTemplate.OrganizationName,
		*params.QueryTemplate.TemplateKey,
	)
	if err != nil {
		log.Println("Error in QueryTemplate:", err.Error())
		return &templateops.PostCertificateTemplateQueryInternalServerError{
			Payload: &apimodels.APIResponse{Message: err.Error()},
		}
	}

	return &templateops.PostCertificateTemplateQueryOK{
		Payload: result,
	}
}

var PostCertificateTemplateQueryHandler = templateops.PostCertificateTemplateQueryHandlerFunc(PostCertificateTemplateQueryImpl)
