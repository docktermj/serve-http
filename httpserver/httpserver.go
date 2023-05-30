package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docktermj/cloudshell/xtermservice"
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
	ApiUrlRoutePrefix              string
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
	SenzingEngineConfigurationJson string
	SenzingModuleName              string
	SenzingVerboseLogging          int
	ServerAddress                  string
	ServerOptions                  []senzinghttpapi.ServerOption
	ServerPort                     int
	SwaggerUrlRoutePrefix          string
	XtermAllowedHostnames          []string
	XtermArguments                 []string
	XtermCommand                   string
	XtermConnectionErrorLimit      int
	XtermKeepalivePingTimeout      int
	XtermMaxBufferSizeBytes        int
	XtermPathLiveness              string
	XtermPathMetrics               string
	XtermPathReadiness             string
	XtermPathXtermjs               string
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
	var userMessage string = ""
	rootMux := http.NewServeMux()

	// Enable Senzing HTTP REST API.

	if httpServer.EnableAll || httpServer.EnableSenzingRestAPI {
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
		srv, err := senzinghttpapi.NewServer(service, httpServer.ServerOptions...)
		if err != nil {
			log.Fatal(err)
		}
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.ApiUrlRoutePrefix), http.StripPrefix("/api", srv))
		userMessage = fmt.Sprintf("%sServing Senzing REST API at http://localhost:%d/%s\n", userMessage, httpServer.ServerPort, httpServer.ApiUrlRoutePrefix)
	}

	// Enable SwaggerUI at /swagger.

	if httpServer.EnableAll || httpServer.EnableSwaggerUI {
		swaggerMux := swaggerui.Handler(httpServer.OpenApiSpecification)
		swaggerFunc := swaggerMux.ServeHTTP
		submux := http.NewServeMux()
		submux.HandleFunc("/", swaggerFunc)
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.SwaggerUrlRoutePrefix), http.StripPrefix("/swagger", submux))
		userMessage = fmt.Sprintf("%sServing SwaggerUI at http://localhost:%d/%s\n", userMessage, httpServer.ServerPort, httpServer.SwaggerUrlRoutePrefix)
	}

	// Enable Xterm at /xterm.

	if httpServer.EnableAll || httpServer.EnableXterm {
		xtermService := &xtermservice.XtermServiceImpl{
			AllowedHostnames:     httpServer.XtermAllowedHostnames,
			Arguments:            httpServer.XtermArguments,
			Command:              httpServer.XtermCommand,
			ConnectionErrorLimit: httpServer.XtermConnectionErrorLimit,
			KeepalivePingTimeout: httpServer.XtermKeepalivePingTimeout,
			MaxBufferSizeBytes:   httpServer.XtermMaxBufferSizeBytes,
			PathLiveness:         httpServer.XtermPathLiveness,
			PathMetrics:          httpServer.XtermPathMetrics,
			PathReadiness:        httpServer.XtermPathReadiness,
			PathXtermjs:          httpServer.XtermPathXtermjs,
			UrlRoutePrefix:       httpServer.XtermUrlRoutePrefix,
		}
		xtermMux := xtermService.Handler(ctx) // Returns *http.ServeMux
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.XtermUrlRoutePrefix), http.StripPrefix("/xterm", xtermMux))
		userMessage = fmt.Sprintf("%sServing XTerm at http://localhost:%d/%s\n", userMessage, httpServer.ServerPort, httpServer.XtermUrlRoutePrefix)
	}

	// Start service.

	if len(userMessage) == 0 {
		userMessage = fmt.Sprintf("Serving on port: %d\n", httpServer.ServerPort)
	}
	fmt.Println(userMessage)
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", httpServer.Port), rootMux); err != nil {
	// 	log.Fatal(err)
	// }
	// return err

	// Start service.

	listenOnAddress := fmt.Sprintf("%s:%v", httpServer.ServerAddress, httpServer.ServerPort)
	server := http.Server{
		Addr:    listenOnAddress,
		Handler: rootMux,
	}
	fmt.Printf("starting server on interface:port '%s'...", listenOnAddress)
	return server.ListenAndServe()
}
