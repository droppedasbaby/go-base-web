package api

import (
	"github.com/go-chi/chi/v5"
)

// restApiRoutes routes the requests to the correct handler.
func restApiRoutes(r chi.Router) {
	// TODO: place api routes here, remove the test route
	r.Route("/v1", func(r chi.Router) {
		r.Get("/test", getTestOperation)
	})
}
