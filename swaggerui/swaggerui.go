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
	Prefix               string
	Port                 int
	OpenApiSpecification []byte
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
	http.Handle(fmt.Sprintf("/%s/", server.Prefix), http.StripPrefix(fmt.Sprintf("/%s", server.Prefix), swaggerui.Handler(server.OpenApiSpecification)))
	fmt.Printf("Serving on port: %d\n", server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil); err != nil {
		log.Fatal(err)
	}
	return err
}
