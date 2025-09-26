package players

import (
	internal "nba-api/internal/database"
	"nba-api/internal/response"
	"nba-api/sql/database"
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
		position := r.URL.Query().Get("position")
		college := r.URL.Query().Get("college")
		yearStartStr := r.URL.Query().Get("year_start")

		var yearStart int32
		if yearStartStr != "" {
			y, err := strconv.Atoi(yearStartStr)
			if err != nil {
				response.ResponseWithError(w, http.StatusBadRequest, "invalid start year format")
				return
			}
			yearStart = int32(y)
		}

		params := database.GetPlayersFilteredParams{
			Column1: position,
			Column2: college,
			Column3: yearStart,
		}

		players, err := s.Queries.GetPlayersFiltered(r.Context(), params)
		if err != nil {
			response.ResponseWithError(w, http.StatusInternalServerError, "error retrieving players")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, players)
	}
}

// GET /v1/players/{id}
func GetPlayerFromIDHandler(s *internal.Store) http.HandlerFunc {
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

// GET /v1/players/search
func GetPlayerFromNameHandler(s *internal.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			response.ResponseWithError(w, http.StatusBadRequest, "no name provided")
			return
		}

		players, err := s.Queries.GetPlayerByName(r.Context(), name)
		if err != nil {
			response.ResponseWithError(w, http.StatusInternalServerError, "errors searching for players")
			return
		}

		response.RespondWithJSON(w, http.StatusOK, players)
	}
}
