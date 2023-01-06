package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenQueryImpl(params tokenops.PostAccessTokenQueryParams) middleware.Responder {
	result, err := tokenapi.QueryToken(
		params.QueryRecord.OrganizationName,
		*params.QueryRecord.TokenID,
	)
	if err != nil {
		return &tokenops.PostAccessTokenQueryInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenQueryOK{
		Payload: result,
	}
}

var PostAccessTokenQueryHandler = tokenops.PostAccessTokenQueryHandlerFunc(PostAccessTokenQueryImpl)
