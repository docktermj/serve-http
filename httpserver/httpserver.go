package httpserver

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
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
	ApiUrlRoutePrefix              string // FIXME: Only works with "api"
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
	SwaggerUrlRoutePrefix          string // FIXME: Only works with "swagger"
	XtermAllowedHostnames          []string
	XtermArguments                 []string
	XtermCommand                   string
	XtermConnectionErrorLimit      int
	XtermKeepalivePingTimeout      int
	XtermMaxBufferSizeBytes        int
	XtermUrlRoutePrefix            string // FIXME: Only works with "xterm"
}

type TemplateVariables struct {
	HttpServerImpl
	ApiServerStatus string
	ApiServerUrl    string
	HtmlTitle       string
	SwaggerStatus   string
	SwaggerUrl      string
	XtermStatus     string
	XtermUrl        string
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

//go:embed static/*
var static embed.FS

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (httpServer *HttpServerImpl) populateStaticTemplate(responseWriter http.ResponseWriter, request *http.Request, filepath string, templateVariables TemplateVariables) {
	templateBytes, err := static.ReadFile(filepath)
	if err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	}
	templateParsed, err := template.New("HtmlTemplate").Parse(string(templateBytes))
	if err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	}
	err = templateParsed.Execute(responseWriter, templateVariables)
	if err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	}
}

func (httpServer *HttpServerImpl) getServerStatus(up bool) string {
	result := "red"
	if httpServer.EnableAll {
		result = "green"
	}
	if up {
		result = "green"
	}
	return result
}

func (httpServer *HttpServerImpl) getServerUrl(up bool, url string) string {
	result := ""
	if httpServer.EnableAll {
		result = url
	}
	if up {
		result = url
	}
	return result
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
	listenOnAddress := fmt.Sprintf("%s:%v", httpServer.ServerAddress, httpServer.ServerPort)
	var userMessage = fmt.Sprintf("Starting server on interface:port '%s'\nServing on http://localhost:%d\n\n", listenOnAddress, httpServer.ServerPort)

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
			UrlRoutePrefix:       httpServer.XtermUrlRoutePrefix,
		}
		xtermMux := xtermService.Handler(ctx) // Returns *http.ServeMux
		rootMux.Handle(fmt.Sprintf("/%s/", httpServer.XtermUrlRoutePrefix), http.StripPrefix("/xterm", xtermMux))
		userMessage = fmt.Sprintf("%sServing XTerm at http://localhost:%d/%s\n", userMessage, httpServer.ServerPort, httpServer.XtermUrlRoutePrefix)
	}

	// Add routes for template pages.

	rootMux.HandleFunc("/overview.html", func(w http.ResponseWriter, r *http.Request) {
		templateVariables := TemplateVariables{
			HttpServerImpl:  *httpServer,
			HtmlTitle:       "Senzing Tools",
			ApiServerUrl:    httpServer.getServerUrl(httpServer.EnableSenzingRestAPI, fmt.Sprintf("http://%s/api", r.Host)),
			ApiServerStatus: httpServer.getServerStatus(httpServer.EnableSenzingRestAPI),
			SwaggerUrl:      httpServer.getServerUrl(httpServer.EnableSwaggerUI, fmt.Sprintf("http://%s/swagger", r.Host)),
			SwaggerStatus:   httpServer.getServerStatus(httpServer.EnableSwaggerUI),
			XtermUrl:        httpServer.getServerUrl(httpServer.EnableXterm, fmt.Sprintf("http://%s/xterm", r.Host)),
			XtermStatus:     httpServer.getServerStatus(httpServer.EnableXterm),
		}
		w.Header().Set("Content-Type", "text/html")
		httpServer.populateStaticTemplate(w, r, "static/templates/overview.html", templateVariables)
	})

	// Add route to static files.

	rootDir, err := fs.Sub(static, "static/root")
	if err != nil {
		panic(err)
	}
	rootMux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(rootDir))))

	// Start service.

	fmt.Println(userMessage)
	server := http.Server{
		Addr:    listenOnAddress,
		Handler: rootMux,
	}
	return server.ListenAndServe()
}
