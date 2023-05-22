package httpserver

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The HttpServer interface...
type HttpServer interface {
	Serve(ctx context.Context) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6204xxxx".
const ComponentId = 9999

// Log message prefix.
const Prefix = "serve-http.httpserver."

// Default gRPC Observer port
const DefaultGrpcObserverPort = "8260"

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates.
var IdMessages = map[int]string{
	2000: "Entry: %+v",
	2001: "SENZING_ENGINE_CONFIGURATION_JSON: %v",
	2002: "Enabling all services",
	2003: "Server listening at %v",
	4001: "Call to net.Listen(tcp, %s) failed.",
	4002: "Call to G2engine.PurgeRepository() failed.",
	4003: "Call to G2engine.Destroy() failed.",
	5001: "Failed to serve.",
}

// Status strings for specific messages.
var IdStatuses = map[int]string{}
