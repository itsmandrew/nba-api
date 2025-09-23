package router

import (
	"net/http"

	internal "nba-api/internal/database"
	md "nba-api/internal/middleware"

	p "nba-api/internal/handlers/players"
	h "nba-api/internal/handlers/utils"

	"github.com/go-chi/chi"
)

func InitRouter(s *internal.Store) http.Handler {
	router := chi.NewRouter()
	router.Use(md.Logger)

	v1Router := chi.NewRouter()
	registerUtilRoutes(v1Router)
	registerPlayerRoutes(v1Router, s)

	router.Mount("/v1", v1Router)

	return router
}

func registerUtilRoutes(r chi.Router) {
	r.Get("/hello", h.HelloHandler)
	r.Get("/health", h.HandlerReady)
	r.Get("/err", h.HandlerErr)
}

func registerPlayerRoutes(r chi.Router, s *internal.Store) {
	r.Get("/lebron", p.GetLeBronHandler(s))
}
