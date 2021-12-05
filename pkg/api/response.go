package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseDTO struct {
	Code    int      `json:"code"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// RespondWithError ...
func RespondWithError(rw http.ResponseWriter, code int, message string) {
	RespondWithJSON(rw, code, ErrorResponseDTO{Code: code, Status: "Error", Message: message})
}

// RespondWithJSON write json
func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(response)
}
