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
	EnableAll                      bool
	EnableSenzingRestAPI           bool
	EnableSwaggerUI                bool
	EnableXterm                    bool
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
	XtermUrlRoutePrefix            string
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
	var err error = nil
	var userMessage string = ""
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

	// Enable Senzing HTTP REST API.

	if httpServer.EnableAll || httpServer.EnableSenzingRestAPI {
		srv, err := senzinghttpapi.NewServer(service, httpServer.ServerOptions...)
		if err != nil {
			log.Fatal(err)
		}
		rootMux.HandleFunc("/", srv.ServeHTTP)
		userMessage = fmt.Sprintf("%sServing Senzing REST API at http://localhost:%d/\n", userMessage, httpServer.Port)
	}

	// Enable SwaggerUI.

	if httpServer.EnableAll || httpServer.EnableSwaggerUI {
		swaggerMux := swaggerui.Handler(httpServer.OpenApiSpecification)
		swaggerFunc := swaggerMux.ServeHTTP
		submux := http.NewServeMux()
		submux.HandleFunc("/", swaggerFunc)
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.SwaggerUrlRoutePrefix), http.StripPrefix(fmt.Sprintf("/%s", httpServer.SwaggerUrlRoutePrefix), submux))
		userMessage = fmt.Sprintf("%sServing SwaggerUI at http://localhost:%d/%s\n", userMessage, httpServer.Port, httpServer.SwaggerUrlRoutePrefix)
	}

	// Enable Xterm.

	if httpServer.EnableAll || httpServer.EnableXterm {
		userMessage = fmt.Sprintf("%sServing XTerm at http://localhost:%d/%s\n", userMessage, httpServer.Port, httpServer.XtermUrlRoutePrefix)

	}

	if len(userMessage) == 0 {
		userMessage = fmt.Sprintf("Serving on port: %d\n", httpServer.Port)
	}

	// Start service.

	fmt.Println(userMessage)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpServer.Port), rootMux); err != nil {
		log.Fatal(err)
	}
	return err
}
