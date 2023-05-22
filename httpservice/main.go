package httpservice

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

// An example constant.
const ExampleConstant = 1

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// An example variable.
var ExampleVariable = map[int]string{
	1: "Just a string",
}
