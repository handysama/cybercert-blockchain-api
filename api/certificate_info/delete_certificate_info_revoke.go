package certificateinfo

import (
	certinfoapi "cybercert-blockchain-api/chaincodeapi/certificate_info"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
)

func DeleteCertificateInfoRevokeImpl(params certinfoops.DeleteCertificateInfoRevokeParams) middleware.Responder {
	result, err := certinfoapi.RevokeCertificate(
		params.RevokeCertificate.OrganizationName,
		*params.RevokeCertificate.CertificateID,
	)
	if err != nil {
		return &certinfoops.DeleteCertificateInfoRevokeInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &certinfoops.DeleteCertificateInfoRevokeOK{
		Payload: &apimodels.APIResponse{
			Message: result.TransactionID,
		},
	}
}

var DeleteCertificateInfoRevokeHandler = certinfoops.DeleteCertificateInfoRevokeHandlerFunc(DeleteCertificateInfoRevokeImpl)
