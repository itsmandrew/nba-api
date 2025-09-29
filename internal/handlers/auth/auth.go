package auth

import (
	"fmt"
	"nba-api/internal/middleware"
	"nba-api/internal/response"
	"net/http"
)

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-Key")
	if apiKey != "API_KEY" {
		response.ResponseWithError(w, http.StatusUnauthorized, "invalid api key")
		return
	}

	token, err := middleware.GenerateJWT("local testing")
	if err != nil {

		response.ResponseWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("failed to generate token: %v", err))
		return
	}

	response.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
