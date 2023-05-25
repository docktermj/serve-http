package swaggerui

import "context"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The HttpServer interface...
type SwaggerUiServer interface {
	Serve(ctx context.Context) error
}
