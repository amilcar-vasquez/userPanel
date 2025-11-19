package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response structure
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// SuccessResponse represents a success response structure
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// RespondJSON sends a JSON response
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, ErrorResponse{
		Success: false,
		Message: message,
	})
}

// RespondSuccess sends a success response
func RespondSuccess(w http.ResponseWriter, data interface{}) {
	RespondJSON(w, http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

// RespondSuccessWithMessage sends a success response with a message
func RespondSuccessWithMessage(w http.ResponseWriter, message string) {
	RespondJSON(w, http.StatusOK, SuccessResponse{
		Success: true,
		Message: message,
	})
}

// ParseJSON parses JSON request body into the provided struct
func ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
