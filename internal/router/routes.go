package router

import (
	"net/http"

	md "nba-api/internal/middleware"

	"github.com/go-chi/chi"
)

func InitRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(md.Logger)

	v1Router := chi.NewRouter()

	router.Mount("/v1", v1Router)

	return router
}
