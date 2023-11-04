package api

import (
	"github.com/go-chi/chi/v5"
)

// restAPIRoutes routes the requests to the correct handler.
func restAPIRoutes(r chi.Router) {
	// TODO: place api routes here, remove the test route
	r.Route("/v1", func(r chi.Router) {
		r.Get("/test", getTestOperation)
	})
}
