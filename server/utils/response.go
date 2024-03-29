package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": message,
		},
	})
}
