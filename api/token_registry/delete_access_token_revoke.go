package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func DeleteAccessTokenRevokeImpl(params tokenops.DeleteAccessTokenRevokeParams) middleware.Responder {
	result, err := tokenapi.RevokeToken(
		*params.TokenRecord.OrganizationName,
		*params.TokenRecord.TokenID,
	)
	if err != nil {
		return &tokenops.DeleteAccessTokenRevokeInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.DeleteAccessTokenRevokeOK{
		Payload: &apimodels.APIResponse{
			Message: result.TransactionID,
		},
	}
}

var DeleteAccessTokenRevokeHandler = tokenops.DeleteAccessTokenRevokeHandlerFunc(DeleteAccessTokenRevokeImpl)
