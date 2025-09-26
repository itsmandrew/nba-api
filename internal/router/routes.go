package router

import (
	"net/http"

	internal "nba-api/internal/database"
	md "nba-api/internal/middleware"

	p "nba-api/internal/handlers/players"
	h "nba-api/internal/handlers/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRouter(s *internal.Store) http.Handler {
	router := chi.NewRouter()

	// Chaining my middleware
	router.Use(md.Logger)
	router.Use(middleware.Recoverer)
	router.Use(md.BearerAuth)

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
	r.Get("/players", p.GetPlayersHandler(s))
	r.Get("/players/{id}", p.GetPlayerFromIDHandler(s))
	r.Get("/players/search", p.GetPlayerFromNameHandler(s))
	r.Get("/players/random", p.GetRandomPlayerHandler(s))
}
