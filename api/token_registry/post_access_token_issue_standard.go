package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenIssueStandardImpl(params tokenops.PostAccessTokenIssueStandardParams) middleware.Responder {
	result, err := tokenapi.IssueStandardToken(
		params.TokenRecord.OrganizationName,
		params.TokenRecord.Args.TokenID,
		params.TokenRecord.Args.IssuerTokenID,
		params.TokenRecord.Args.Recipient,
		params.TokenRecord.Args.Amount,
		params.TokenRecord.Args.AccessQuota,
		params.TokenRecord.Args.ExpiryDate,
	)
	if err != nil {
		return &tokenops.PostAccessTokenIssueStandardInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenIssueStandardCreated{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostAccessTokenIssueStandardHandler = tokenops.PostAccessTokenIssueStandardHandlerFunc(PostAccessTokenIssueStandardImpl)
