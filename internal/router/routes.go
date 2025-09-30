package router

import (
	"net/http"
	"time"

	internal "nba-api/internal/database"
	md "nba-api/internal/middleware"

	a "nba-api/internal/handlers/auth"
	p "nba-api/internal/handlers/players"
	h "nba-api/internal/handlers/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
)

func InitRouter(s *internal.Store) http.Handler {
	router := chi.NewRouter()

	// Chaining my middleware
	router.Use(md.Logger)
	router.Use(middleware.Recoverer)

	router.Use(httprate.LimitByIP(20, time.Minute))

	router.Route("/v1", func(r chi.Router) {
		// ---- Public Routes -----
		registerAuthRoutes(r)
		registerUtilRoutes(r)

		// ---- Private Routes ----
		r.Group(func(private chi.Router) {
			private.Use(md.JWTAuth)

			private.Use(httprate.LimitByIP(20, time.Minute))

			registerPlayerRoutes(private, s)
			registerOtherRoutes(private, s)
		})
	})

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

func registerAuthRoutes(r chi.Router) {
	r.Post("/generate-token", a.GenerateTokenHandler)
}

func registerOtherRoutes(r chi.Router, s *internal.Store) {
	r.Get("/colleges", p.GetAllCollegesHander(s))
}
