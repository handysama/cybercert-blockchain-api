// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	certificateinfoapi "cybercert-blockchain-api/api/certificate_info"
	certificatetemplateapi "cybercert-blockchain-api/api/certificate_template"
	healthcheckapi "cybercert-blockchain-api/api/healthcheck"
	accesstokenapi "cybercert-blockchain-api/api/token_registry"

	"cybercert-blockchain-api/restapi/operations"
	"cybercert-blockchain-api/restapi/operations/certificate_info"
	"cybercert-blockchain-api/restapi/operations/certificate_template"
	"cybercert-blockchain-api/restapi/operations/healthcheck"
	"cybercert-blockchain-api/restapi/operations/token_registry"
)

//go:generate swagger generate server --target ../../cybercert-blockchain-api --name CybercertBlockchainAPI --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.CybercertBlockchainAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CybercertBlockchainAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CertificateInfoDeleteCertificateInfoRevokeHandler = certificateinfoapi.DeleteCertificateInfoRevokeHandler
	api.CertificateInfoPostCertificateInfoHistoryHandler = certificateinfoapi.PostCertificateInfoHistoryHandler
	api.CertificateInfoPostCertificateInfoIssueHandler = certificateinfoapi.PostCertificateInfoIssueHandler
	api.CertificateInfoPostCertificateInfoIssueBatchHandler = certificateinfoapi.PostCertificateInfoIssueBatchHandler
	api.CertificateInfoPostCertificateInfoQueryCertificateHandler = certificateinfoapi.PostCertificateInfoQueryCertificateHandler
	api.CertificateInfoPostCertificateInfoQueryRecordsHandler = certificateinfoapi.PostCertificateInfoQueryRecordsHandler

	api.CertificateTemplatePostCertificateTemplateHandler = certificatetemplateapi.PostCertificateTemplateHandler
	api.CertificateTemplatePostCertificateTemplateQueryHandler = certificatetemplateapi.PostCertificateTemplateQueryHandler
	api.CertificateTemplatePostCertificateTemplateHistoryHandler = certificatetemplateapi.PostCertificateTemplateHistoryHandler

	api.TokenRegistryDeleteAccessTokenRevokeHandler = accesstokenapi.DeleteAccessTokenRevokeHandler
	api.TokenRegistryPostAccessTokenChangeTokenOwnerHandler = accesstokenapi.PostAccessTokenChangeTokenOwnerHandler
	api.TokenRegistryPostAccessTokenConsumeHandler = accesstokenapi.PostAccessTokenConsumeHandler
	api.TokenRegistryPostAccessTokenHistoryHandler = accesstokenapi.PostAccessTokenHistoryHandler
	api.TokenRegistryPostAccessTokenIssueRootHandler = accesstokenapi.PostAccessTokenIssueRootHandler
	api.TokenRegistryPostAccessTokenIssueTransferableHandler = accesstokenapi.PostAccessTokenIssueTransferableHandler
	api.TokenRegistryPostAccessTokenIssueStandardHandler = accesstokenapi.PostAccessTokenIssueStandardHandler
	api.TokenRegistryPostAccessTokenQueryHandler = accesstokenapi.PostAccessTokenQueryHandler
	api.TokenRegistryPostAccessTokenQueryRecordsHandler = accesstokenapi.PostAccessTokenQueryRecordsHandler
	api.TokenRegistryPostAccessTokenStatusHandler = accesstokenapi.PostAccessTokenStatusHandler

	api.HealthcheckGetHealthcheckReadyHandler = healthcheckapi.GetHealthcheckReadyHandler

	if api.TokenRegistryDeleteAccessTokenRevokeHandler == nil {
		api.TokenRegistryDeleteAccessTokenRevokeHandler = token_registry.DeleteAccessTokenRevokeHandlerFunc(func(params token_registry.DeleteAccessTokenRevokeParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.DeleteAccessTokenRevoke has not yet been implemented")
		})
	}
	if api.CertificateInfoDeleteCertificateInfoRevokeHandler == nil {
		api.CertificateInfoDeleteCertificateInfoRevokeHandler = certificate_info.DeleteCertificateInfoRevokeHandlerFunc(func(params certificate_info.DeleteCertificateInfoRevokeParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.DeleteCertificateInfoRevoke has not yet been implemented")
		})
	}
	if api.HealthcheckGetHealthcheckReadyHandler == nil {
		api.HealthcheckGetHealthcheckReadyHandler = healthcheck.GetHealthcheckReadyHandlerFunc(func(params healthcheck.GetHealthcheckReadyParams) middleware.Responder {
			return middleware.NotImplemented("operation healthcheck.GetHealthcheckReady has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenChangeTokenOwnerHandler == nil {
		api.TokenRegistryPostAccessTokenChangeTokenOwnerHandler = token_registry.PostAccessTokenChangeTokenOwnerHandlerFunc(func(params token_registry.PostAccessTokenChangeTokenOwnerParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenChangeTokenOwner has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenConsumeHandler == nil {
		api.TokenRegistryPostAccessTokenConsumeHandler = token_registry.PostAccessTokenConsumeHandlerFunc(func(params token_registry.PostAccessTokenConsumeParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenConsume has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenHistoryHandler == nil {
		api.TokenRegistryPostAccessTokenHistoryHandler = token_registry.PostAccessTokenHistoryHandlerFunc(func(params token_registry.PostAccessTokenHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenHistory has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenIssueRootHandler == nil {
		api.TokenRegistryPostAccessTokenIssueRootHandler = token_registry.PostAccessTokenIssueRootHandlerFunc(func(params token_registry.PostAccessTokenIssueRootParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenIssueRoot has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenIssueStandardHandler == nil {
		api.TokenRegistryPostAccessTokenIssueStandardHandler = token_registry.PostAccessTokenIssueStandardHandlerFunc(func(params token_registry.PostAccessTokenIssueStandardParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenIssueStandard has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenIssueTransferableHandler == nil {
		api.TokenRegistryPostAccessTokenIssueTransferableHandler = token_registry.PostAccessTokenIssueTransferableHandlerFunc(func(params token_registry.PostAccessTokenIssueTransferableParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenIssueTransferable has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenQueryHandler == nil {
		api.TokenRegistryPostAccessTokenQueryHandler = token_registry.PostAccessTokenQueryHandlerFunc(func(params token_registry.PostAccessTokenQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenQuery has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenQueryRecordsHandler == nil {
		api.TokenRegistryPostAccessTokenQueryRecordsHandler = token_registry.PostAccessTokenQueryRecordsHandlerFunc(func(params token_registry.PostAccessTokenQueryRecordsParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenQueryRecords has not yet been implemented")
		})
	}
	if api.TokenRegistryPostAccessTokenStatusHandler == nil {
		api.TokenRegistryPostAccessTokenStatusHandler = token_registry.PostAccessTokenStatusHandlerFunc(func(params token_registry.PostAccessTokenStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation token_registry.PostAccessTokenStatus has not yet been implemented")
		})
	}
	if api.CertificateInfoPostCertificateInfoHistoryHandler == nil {
		api.CertificateInfoPostCertificateInfoHistoryHandler = certificate_info.PostCertificateInfoHistoryHandlerFunc(func(params certificate_info.PostCertificateInfoHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.PostCertificateInfoHistory has not yet been implemented")
		})
	}
	if api.CertificateInfoPostCertificateInfoIssueHandler == nil {
		api.CertificateInfoPostCertificateInfoIssueHandler = certificate_info.PostCertificateInfoIssueHandlerFunc(func(params certificate_info.PostCertificateInfoIssueParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.PostCertificateInfoIssue has not yet been implemented")
		})
	}
	if api.CertificateInfoPostCertificateInfoIssueBatchHandler == nil {
		api.CertificateInfoPostCertificateInfoIssueBatchHandler = certificate_info.PostCertificateInfoIssueBatchHandlerFunc(func(params certificate_info.PostCertificateInfoIssueBatchParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.PostCertificateInfoIssueBatch has not yet been implemented")
		})
	}
	if api.CertificateInfoPostCertificateInfoQueryCertificateHandler == nil {
		api.CertificateInfoPostCertificateInfoQueryCertificateHandler = certificate_info.PostCertificateInfoQueryCertificateHandlerFunc(func(params certificate_info.PostCertificateInfoQueryCertificateParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.PostCertificateInfoQueryCertificate has not yet been implemented")
		})
	}
	if api.CertificateInfoPostCertificateInfoQueryRecordsHandler == nil {
		api.CertificateInfoPostCertificateInfoQueryRecordsHandler = certificate_info.PostCertificateInfoQueryRecordsHandlerFunc(func(params certificate_info.PostCertificateInfoQueryRecordsParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_info.PostCertificateInfoQueryRecords has not yet been implemented")
		})
	}
	if api.CertificateTemplatePostCertificateTemplateHandler == nil {
		api.CertificateTemplatePostCertificateTemplateHandler = certificate_template.PostCertificateTemplateHandlerFunc(func(params certificate_template.PostCertificateTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_template.PostCertificateTemplate has not yet been implemented")
		})
	}
	if api.CertificateTemplatePostCertificateTemplateHistoryHandler == nil {
		api.CertificateTemplatePostCertificateTemplateHistoryHandler = certificate_template.PostCertificateTemplateHistoryHandlerFunc(func(params certificate_template.PostCertificateTemplateHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_template.PostCertificateTemplateHistory has not yet been implemented")
		})
	}
	if api.CertificateTemplatePostCertificateTemplateQueryHandler == nil {
		api.CertificateTemplatePostCertificateTemplateQueryHandler = certificate_template.PostCertificateTemplateQueryHandlerFunc(func(params certificate_template.PostCertificateTemplateQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation certificate_template.PostCertificateTemplateQuery has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	// Read config
	viper.AutomaticEnv()
	viper.AddConfigPath("config/")
	viper.ReadInConfig()

	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
