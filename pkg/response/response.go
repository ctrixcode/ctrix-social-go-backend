package response

import (
	"encoding/json"
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

// JSONSuccess sends a success JSON response.
func JSONSuccess(w http.ResponseWriter, data interface{}, statusCode int, message ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := SuccessResponse{
		Success: true,
		Data:    data,
	}
	if len(message) > 0 {
		resp.Message = message[0]
	}

	json.NewEncoder(w).Encode(resp)
}

func JSONError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	apiErr, ok := err.(*errors.APIError)
	if !ok {
		apiErr = errors.InternalServerError(errors.ErrInternalServerError)
	}

	w.WriteHeader(apiErr.StatusCode)

	resp := ErrorResponse{
		Success: false,
		Code:    apiErr.Type,
		Message: apiErr.Message,
		Details: apiErr.Details,
	}

	json.NewEncoder(w).Encode(resp)
}
