//go:build go1.20

package httpbase

import "net/http"

// ResponseController is type net/http.ResponseController which was added in Go 1.20.

type ResponseController = http.ResponseController
