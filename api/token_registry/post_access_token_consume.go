package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenConsumeImpl(params tokenops.PostAccessTokenConsumeParams) middleware.Responder {
	result, err := tokenapi.ConsumeToken(
		params.TokenRecord.OrganizationName,
		*params.TokenRecord.TokenID,
	)
	if err != nil {
		return &tokenops.PostAccessTokenConsumeInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenConsumeOK{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostAccessTokenConsumeHandler = tokenops.PostAccessTokenConsumeHandlerFunc(PostAccessTokenConsumeImpl)
