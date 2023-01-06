package certificateinfo

import (
	certinfoapi "cybercert-blockchain-api/chaincodeapi/certificate_info"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateInfoQueryCertificateImpl(params certinfoops.PostCertificateInfoQueryCertificateParams) middleware.Responder {
	result, err := certinfoapi.QueryCertificate(
		params.QueryCertificate.OrganizationName,
		*params.QueryCertificate.CertificateID,
	)
	if err != nil {
		return &certinfoops.PostCertificateInfoQueryCertificateInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &certinfoops.PostCertificateInfoQueryCertificateOK{
		Payload: result,
	}
}

var PostCertificateInfoQueryCertificateHandler = certinfoops.PostCertificateInfoQueryCertificateHandlerFunc(PostCertificateInfoQueryCertificateImpl)
