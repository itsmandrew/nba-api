package router

import (
	"net/http"

	md "nba-api/internal/middleware"

	h "nba-api/internal/handlers/utils"

	"github.com/go-chi/chi"
)

func InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(md.Logger)

	v1Router := chi.NewRouter()
	registerUtilRoutes(v1Router)

	router.Mount("/v1", v1Router)

	return router
}

func registerUtilRoutes(r chi.Router) {
	r.Get("/hello", h.HelloHandler)
	r.Get("/health", h.HandlerReady)
	r.Get("/err", h.HandlerErr)
}
