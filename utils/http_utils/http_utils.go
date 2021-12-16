package http_utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// func ResponsdError(w http.ResponseWriter, err restErr.RestErr) {
// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(err.Status)
// 	// json.NewEncoder(w).Encode(err)
// 	RespondJson(w, err.Status, err)
// }
