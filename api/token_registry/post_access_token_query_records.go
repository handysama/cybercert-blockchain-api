package tokenregistry

import (
	tokenapi "cybercert-blockchain-api/chaincodeapi/token_registry"
	apimodels "cybercert-blockchain-api/models"
	tokenops "cybercert-blockchain-api/restapi/operations/token_registry"

	"github.com/go-openapi/runtime/middleware"
)

func PostAccessTokenQueryRecordsImpl(params tokenops.PostAccessTokenQueryRecordsParams) middleware.Responder {
	result, err := tokenapi.QueryRecords(
		params.QueryRecords.OrganizationName,
		*params.QueryRecords.QueryString,
	)
	if err != nil {
		return &tokenops.PostAccessTokenQueryRecordsInternalServerError{
			Payload: &apimodels.APIResponse{
				Message: err.Error(),
			},
		}
	}

	return &tokenops.PostAccessTokenQueryRecordsOK{
		Payload: result,
	}
}

var PostAccessTokenQueryRecordsHandler = tokenops.PostAccessTokenQueryRecordsHandlerFunc(PostAccessTokenQueryRecordsImpl)
