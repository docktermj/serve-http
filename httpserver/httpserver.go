package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docktermj/go-http/senzinghttpapi"
	"github.com/docktermj/serve-http/httpservice"
	"github.com/flowchartsman/swaggerui"
	"github.com/senzing/go-observing/observer"
	"google.golang.org/grpc"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServerImpl is the default implementation of the HttpServer interface.
type HttpServerImpl struct {
	EnableSwaggerUI                bool
	GrpcDialOptions                []grpc.DialOption
	GrpcTarget                     string
	LogLevelName                   string
	ObserverOrigin                 string
	Observers                      []observer.Observer
	OpenApiSpecification           []byte
	Port                           int
	SenzingEngineConfigurationJson string
	SenzingModuleName              string
	SenzingVerboseLogging          int
	ServerOptions                  []senzinghttpapi.ServerOption
	SwaggerUrlRoutePrefix          string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Serve method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/

func (httpServer *HttpServerImpl) Serve(ctx context.Context) error {
	rootMux := http.NewServeMux()

	// Create Senzing HTTP REST API service instance.

	service := &httpservice.HttpServiceImpl{
		GrpcDialOptions:                httpServer.GrpcDialOptions,
		GrpcTarget:                     httpServer.GrpcTarget,
		LogLevelName:                   httpServer.LogLevelName,
		ObserverOrigin:                 httpServer.ObserverOrigin,
		Observers:                      httpServer.Observers,
		SenzingEngineConfigurationJson: httpServer.SenzingEngineConfigurationJson,
		SenzingModuleName:              httpServer.SenzingModuleName,
		SenzingVerboseLogging:          httpServer.SenzingVerboseLogging,
	}

	// Create generated server.

	srv, err := senzinghttpapi.NewServer(service, httpServer.ServerOptions...)
	if err != nil {
		log.Fatal(err)
	}

	// Enable Senzing HTTP REST API.

	rootMux.HandleFunc("/", srv.ServeHTTP)
	userMessage := fmt.Sprintf("Serving Senzing REST API on port: %d\n", httpServer.Port)

	// Optionally enable SwaggerUI.

	if httpServer.EnableSwaggerUI {
		swaggerMux := swaggerui.Handler(httpServer.OpenApiSpecification)
		swaggerFunc := swaggerMux.ServeHTTP
		submux := http.NewServeMux()
		submux.HandleFunc("/", swaggerFunc)
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.SwaggerUrlRoutePrefix), http.StripPrefix(fmt.Sprintf("/%s", httpServer.SwaggerUrlRoutePrefix), submux))
		userMessage = fmt.Sprintf("%sView SwaggerUI at http://localhost:%d/%s\n", userMessage, httpServer.Port, httpServer.SwaggerUrlRoutePrefix)
	}

	// Start service.

	fmt.Println(userMessage)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpServer.Port), rootMux); err != nil {
		log.Fatal(err)
	}
	return err
}
