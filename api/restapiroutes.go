package api

import (
	"github.com/go-chi/chi/v5"
)

// restApiRoutes routes the requests to the correct handler.
func restApiRoutes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {})
}
