package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenChangeTokenOwnerImpl(params tokenops.PostAccessTokenChangeTokenOwnerParams) middleware.Responder {
	result, err := tokenapi.ChangeTokenOwner(
		params.TokenRecord.OrganizationName,
		*params.TokenRecord.TokenID,
		*params.TokenRecord.Owner,
	)
	if err != nil {
		return &tokenops.PostAccessTokenChangeTokenOwnerInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenChangeTokenOwnerOK{
		Payload: &apimodels.CreatedResponse{
			TransactionID: result.TransactionID,
		},
	}
}

var PostAccessTokenChangeTokenOwnerHandler = tokenops.PostAccessTokenChangeTokenOwnerHandlerFunc(PostAccessTokenChangeTokenOwnerImpl)
