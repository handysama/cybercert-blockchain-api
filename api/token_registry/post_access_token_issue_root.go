package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenIssueRootImpl(params tokenops.PostAccessTokenIssueRootParams) middleware.Responder {
	result, err := tokenapi.IssueRootToken(
		params.TokenRecord.OrganizationName,
		params.TokenRecord.Args.TokenID,
		params.TokenRecord.Args.CertificateID,
		params.TokenRecord.Args.Owner,
	)
	if err != nil {
		return &tokenops.PostAccessTokenIssueRootInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenIssueRootCreated{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostAccessTokenIssueRootHandler = tokenops.PostAccessTokenIssueRootHandlerFunc(PostAccessTokenIssueRootImpl)
