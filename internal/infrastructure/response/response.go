package response

import (
	"encoding/json"
	"net/http"
)

// Response represents the structure of the response to send to the client
type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// BuildResponse is a utility function to build and send a response
func BuildResponse(w http.ResponseWriter, statusCode int, data Response) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
