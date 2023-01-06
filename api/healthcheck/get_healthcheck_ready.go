package healthcheck

import (
	healthcheckops "cybercert-blockchain-api/restapi/operations/healthcheck"

	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/viper"
)

func GetHealthcheckReadyImpl(params healthcheckops.GetHealthcheckReadyParams) middleware.Responder {
	return &healthcheckops.GetHealthcheckReadyOK{
		Payload: &healthcheckops.GetHealthcheckReadyOKBody{
			Message: "OK",
			Version: viper.GetString("APP_VERSION_HASH"),
		},
	}
}

var GetHealthcheckReadyHandler = healthcheckops.GetHealthcheckReadyHandlerFunc(GetHealthcheckReadyImpl)
