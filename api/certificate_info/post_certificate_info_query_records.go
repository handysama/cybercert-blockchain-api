package certificateinfo

import (
	certinfoapi "cybercert-blockchain-api/chaincodeapi/certificate_info"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateInfoQueryRecordsImpl(params certinfoops.PostCertificateInfoQueryRecordsParams) middleware.Responder {
	result, err := certinfoapi.QueryRecords(
		params.QueryRecords.OrganizationName,
		*params.QueryRecords.QueryString,
	)
	if err != nil {
		return &certinfoops.PostCertificateInfoQueryRecordsInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &certinfoops.PostCertificateInfoQueryRecordsOK{
		Payload: result,
	}
}

var PostCertificateInfoQueryRecordsHandler = certinfoops.PostCertificateInfoQueryRecordsHandlerFunc(PostCertificateInfoQueryRecordsImpl)
