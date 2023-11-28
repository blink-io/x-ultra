package roundtripper

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"
)

type Func func(*http.Request) (*http.Response, error)

func (f Func) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
