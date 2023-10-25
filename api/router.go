package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-base-web/utils"
	"go.uber.org/zap"
	"net/http"
)

// GetRouter returns a chi router with the correct routes.
// Both for swagger router and normal api routes
func GetRouter(logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	if logger != nil {
		r.Use(utils.Logger(logger))
	}
	r.Use(middleware.Recoverer)

	buildTopLevelRoutes(r)

	return r
}

// buildTopLevelRoutes builds the top level routes for the api.
// This includes the swagger routes, and the basic api routes.
func buildTopLevelRoutes(r *chi.Mux) {
	r.Get("/docs/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/docs", http.FileServer(http.Dir("./docs"))).ServeHTTP(w, r)
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
	))

	r.Route("/api", restApiRoutes)
}
