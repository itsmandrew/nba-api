package players

import (
	internal "nba-api/internal/database"
	"nba-api/internal/response"
	"net/http"
)

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
