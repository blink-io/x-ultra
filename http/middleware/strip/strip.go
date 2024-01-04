package strip

import (
	chimw "github.com/go-chi/chi/v5/middleware"
)

var NewHandler = chimw.StripSlashes
