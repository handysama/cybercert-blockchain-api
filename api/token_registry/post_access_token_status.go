package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenStatusImpl(params tokenops.PostAccessTokenStatusParams) middleware.Responder {
	tokenId := *params.TokenRecord.TokenID
	tokenStatus, err := tokenapi.QueryTokenStatus(
		params.TokenRecord.OrganizationName,
		tokenId,
	)
	if err != nil {
		return &tokenops.PostAccessTokenStatusInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenStatusOK{
		Payload: &tokenops.PostAccessTokenStatusOKBody{
			TokenID: tokenId,
			Status:  tokenStatus,
		},
	}
}

var PostAccessTokenStatusHandler = tokenops.PostAccessTokenStatusHandlerFunc(PostAccessTokenStatusImpl)
