package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/docktermj/go-http/senzinghttpapi"
	"github.com/docktermj/serve-http/httpservice"
	"github.com/ogen-go/ogen/middleware"
	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/logging"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/observerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServerImpl is the default implementation of the HttpServer interface.
type HttpServerImpl struct {
	GrpcDialOptions                []grpc.DialOption
	GrpcTarget                     string
	logger                         logging.LoggingInterface
	LogLevelName                   string
	ObserverOrigin                 string
	Observers                      []observer.Observer
	ObserverUrl                    string
	Port                           int
	SenzingEngineConfigurationJson string
	SenzingModuleName              string
	SenzingVerboseLogging          int
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const exampleConstant = "examplePackage"

var bobBool bool

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// --- Logging -------------------------------------------------------------------------

// Get the Logger singleton.
func (httpServer *HttpServerImpl) getLogger() logging.LoggingInterface {
	var err error = nil
	if httpServer.logger == nil {
		options := []interface{}{
			&logging.OptionCallerSkip{Value: 3},
		}
		httpServer.logger, err = logging.NewSenzingToolsLogger(ComponentId, IdMessages, options...)
		if err != nil {
			panic(err)
		}
	}
	return httpServer.logger
}

// Log message.
func (httpServer *HttpServerImpl) log(messageNumber int, details ...interface{}) {
	httpServer.getLogger().Log(messageNumber, details...)
}

// --- Observing --------------------------------------------------------------

func (httpServer *HttpServerImpl) createGrpcObserver(ctx context.Context, parsedUrl url.URL) (observer.Observer, error) {
	var err error
	var result observer.Observer

	port := DefaultGrpcObserverPort
	if len(parsedUrl.Port()) > 0 {
		port = parsedUrl.Port()
	}
	target := fmt.Sprintf("%s:%s", parsedUrl.Hostname(), port)

	// TODO: Allow specification of options from ObserverUrl/parsedUrl
	grpcDialOptions := grpc.WithTransportCredentials(insecure.NewCredentials())

	grpcConnection, err := grpc.Dial(target, grpcDialOptions)
	if err != nil {
		return result, err
	}
	result = &observer.ObserverGrpc{
		GrpcClient: observerpb.NewObserverClient(grpcConnection),
		Id:         "serve-http",
	}
	return result, err
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (httpServer *HttpServerImpl) addResponseHeaders() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		fmt.Println(">>>>>> Hi there")
		logger.Info("Handling request")
		resp, err := next(req)
		return resp, err
	}
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The SaySomething method simply prints the 'Something' value in the type-struct.

Input
  - ctx: A context to control lifecycle.

Output
  - Nothing is returned, except for an error.  However, something is printed.
    See the example output.
*/

func (httpServer *HttpServerImpl) Serve(ctx context.Context) error {

	// Create service instance.

	service := &httpservice.HttpServiceImpl{
		GrpcDialOptions:                httpServer.GrpcDialOptions,
		GrpcTarget:                     httpServer.GrpcTarget,
		LogLevelName:                   httpServer.LogLevelName,
		ObserverOrigin:                 httpServer.ObserverOrigin,
		ObserverUrl:                    httpServer.ObserverUrl,
		SenzingEngineConfigurationJson: httpServer.SenzingEngineConfigurationJson,
		SenzingModuleName:              httpServer.SenzingModuleName,
		SenzingVerboseLogging:          httpServer.SenzingVerboseLogging,
	}

	// Create generated server options.

	serverOptions := []senzinghttpapi.ServerOption{
		// httpServer.addResponseHeaders,
	}

	httpServer.addResponseHeaders()

	// Create generated server.

	srv, err := senzinghttpapi.NewServer(service, serverOptions...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serving on port: %d\n", httpServer.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpServer.Port), srv); err != nil {
		log.Fatal(err)
	}

	return err

}
