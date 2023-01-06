package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenHistoryImpl(params tokenops.PostAccessTokenHistoryParams) middleware.Responder {
	result, err := tokenapi.GetHistoryForKey(
		params.TokenRecord.OrganizationName,
		params.TokenRecord.TokenID,
	)
	if err != nil {
		return &tokenops.PostAccessTokenHistoryInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenHistoryOK{
		Payload: result,
	}
}

var PostAccessTokenHistoryHandler = tokenops.PostAccessTokenHistoryHandlerFunc(PostAccessTokenHistoryImpl)
