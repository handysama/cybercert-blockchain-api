package certificateinfo

import (
	"encoding/json"

	certinfoapi "cybercert-blockchain-api/chaincodeapi/certificate_info"
	apimodels "cybercert-blockchain-api/models"
	certinfoops "cybercert-blockchain-api/restapi/operations/certificate_info"

	"github.com/go-openapi/runtime/middleware"
)

func PostCertificateInfoIssueImpl(params certinfoops.PostCertificateInfoIssueParams) middleware.Responder {
	extrasBytes, err := json.Marshal(params.CertificateRecord.CertificateDetails.Extras)
	if err != nil {
		return &certinfoops.PostCertificateInfoIssueInternalServerError{}
	}

	certDetails := params.CertificateRecord.CertificateDetails
	orgName := params.CertificateRecord.OrganizationName
	result, err := certinfoapi.IssueCertificate(
		orgName,
		certDetails.CertificateID,
		certDetails.CertificateSignature,
		certDetails.TemplateRef,
		certDetails.CourseName,
		certDetails.ModuleName,
		certDetails.CertificateHolder,
		certDetails.Email,
		orgName,
		certDetails.IssuerName,
		certDetails.IssuedAt,
		extrasBytes,
	)
	if err != nil {
		return &certinfoops.PostCertificateInfoIssueInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &certinfoops.PostCertificateInfoIssueCreated{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostCertificateInfoIssueHandler = certinfoops.PostCertificateInfoIssueHandlerFunc(PostCertificateInfoIssueImpl)
