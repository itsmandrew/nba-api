package players

import (
	internal "nba-api/internal/database"
	"nba-api/internal/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// GET /v1/lebron
func GetLeBronHandler(s *internal.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		player, err := s.Queries.GetLeBronJames(r.Context())
		if err != nil {
			response.ResponseWithError(w, http.StatusInternalServerError, "error with retrieving lebron")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, player)
	}
}

// GET /v1/players
func GetPlayersHandler(s *internal.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := s.Queries.GetPlayers(r.Context())
		if err != nil {
			response.ResponseWithError(w, http.StatusInternalServerError, "error retrieving players")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, players)
	}
}

// GET /v1/players/{id}
func GetPlayerFromID(s *internal.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.ResponseWithError(w, http.StatusBadRequest, "invalid player id")
			return
		}

		player, err := s.Queries.GetPlayerByID(r.Context(), int32(id))
		if err != nil {
			response.ResponseWithError(w, http.StatusNotFound, "player not found")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, player)
	}
}
