package utils

import (
	"nba-api/internal/response"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}

func HandlerReady(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, struct{}{})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	response.ResponseWithError(w, http.StatusBadRequest, "something went wrong")
}
