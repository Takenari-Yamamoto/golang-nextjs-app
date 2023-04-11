// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Takenari-Yamamoto/golang-nextjs-app/restapi/operations"
)

//go:generate swagger generate server --target ../../backend --name SampleAPI --spec ../../schema/swagger.yml --model-package restapi/models --principal interface{}

func configureFlags(api *operations.SampleAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SampleAPIAPI) http.Handler {
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

	if api.DeleteUsersIDHandler == nil {
		api.DeleteUsersIDHandler = operations.DeleteUsersIDHandlerFunc(func(params operations.DeleteUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteUsersID has not yet been implemented")
		})
	}
	if api.GetUsersHandler == nil {
		api.GetUsersHandler = operations.GetUsersHandlerFunc(func(params operations.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUsers has not yet been implemented")
		})
	}
	if api.GetUsersIDHandler == nil {
		api.GetUsersIDHandler = operations.GetUsersIDHandlerFunc(func(params operations.GetUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUsersID has not yet been implemented")
		})
	}
	if api.PostUsersHandler == nil {
		api.PostUsersHandler = operations.PostUsersHandlerFunc(func(params operations.PostUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUsers has not yet been implemented")
		})
	}
	if api.PutUsersIDHandler == nil {
		api.PutUsersIDHandler = operations.PutUsersIDHandlerFunc(func(params operations.PutUsersIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PutUsersID has not yet been implemented")
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
	return handler
}
