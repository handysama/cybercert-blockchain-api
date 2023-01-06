package certificateinfo

import (
	certinfoapi "cybercert-blockchain-api/chaincodeapi/certificate_info"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateInfoHistoryImpl(params certinfoops.PostCertificateInfoHistoryParams) middleware.Responder {
	result, err := certinfoapi.GetHistoryForKey(
		params.GetHistory.OrganizationName,
		*params.GetHistory.CertificateID,
	)
	if err != nil {
		return &certinfoops.PostCertificateInfoHistoryInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &certinfoops.PostCertificateInfoHistoryOK{
		Payload: result,
	}
}

var PostCertificateInfoHistoryHandler = certinfoops.PostCertificateInfoHistoryHandlerFunc(PostCertificateInfoHistoryImpl)
