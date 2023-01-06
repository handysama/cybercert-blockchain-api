package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenIssueTransferableImpl(params tokenops.PostAccessTokenIssueTransferableParams) middleware.Responder {
	result, err := tokenapi.IssueTransferableToken(
		params.TokenRecord.OrganizationName,
		params.TokenRecord.Args.TokenID,
		params.TokenRecord.Args.IssuerTokenID,
		params.TokenRecord.Args.Recipient,
		params.TokenRecord.Args.Amount,
		params.TokenRecord.Args.MonthlyTokenQuota,
		params.TokenRecord.Args.ExpiryDate,
	)
	if err != nil {
		return &tokenops.PostAccessTokenIssueTransferableInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenIssueTransferableCreated{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostAccessTokenIssueTransferableHandler = tokenops.PostAccessTokenIssueTransferableHandlerFunc(PostAccessTokenIssueTransferableImpl)
