package senzingapiservice

import (
	"github.com/docktermj/go-http/senzinghttpapi"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The HttpService interface is...
type HttpService interface {
	senzinghttpapi.Handler
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6503xxxx".
const ComponentId = 9999

// Log message prefix.
const Prefix = "serve-http.httpservice."

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2config implementations.
var IdMessages = map[int]string{
	10: "Enter " + Prefix + "InitializeSenzing().",
}

// Status strings for specific messages.
var IdStatuses = map[int]string{}
