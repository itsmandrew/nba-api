package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Reading with 5xx error: ", message)
	}

	RespondWithJSON(w, code, errResponse{
		Error: message,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
