package swaggerui

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// HttpServiceImpl is...
type SwaggerUiServerImpl struct {
	SwaggerUrlRoutePrefix string
	Port                  int
	OpenApiSpecification  []byte
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Serve method...

Input
  - ctx: A context to control lifecycle.
*/

func (server *SwaggerUiServerImpl) Serve(ctx context.Context) error {
	var err error = nil
	http.Handle(fmt.Sprintf("/%s/", server.SwaggerUrlRoutePrefix), http.StripPrefix(fmt.Sprintf("/%s", server.SwaggerUrlRoutePrefix), swaggerui.Handler(server.OpenApiSpecification)))
	fmt.Printf("SwaggerUI visible at http://localhost:%d/swagger\n", server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil); err != nil {
		log.Fatal(err)
	}
	return err
}
