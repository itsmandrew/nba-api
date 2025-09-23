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
