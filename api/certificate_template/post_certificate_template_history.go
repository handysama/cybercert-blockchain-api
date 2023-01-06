package certificatetemplate

import (
	cert_template_api "cybercert-blockchain-api/chaincodeapi/certificate_template"
	apimodels "cybercert-blockchain-api/models"
	templateops "cybercert-blockchain-api/restapi/operations/certificate_template"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateTemplateHistoryImpl(params templateops.PostCertificateTemplateHistoryParams) middleware.Responder {
	result, err := cert_template_api.GetHistoryForKey(
		params.GetHistory.OrganizationName,
		*params.GetHistory.TemplateKey,
	)
	if err != nil {
		return &templateops.PostCertificateTemplateHistoryInternalServerError{
			Payload: &apimodels.APIResponse{Message: err.Error()},
		}
	}

	return &templateops.PostCertificateTemplateHistoryOK{
		Payload: result,
	}
}

var PostCertificateTemplateHistoryHandler = templateops.PostCertificateTemplateHistoryHandlerFunc(PostCertificateTemplateHistoryImpl)
