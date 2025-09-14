package errors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Type          string      `json:"errorType"`
	StatusCode    int         `json:"statusCode"`
	Message       string      `json:"message,omitempty"`
	Details       interface{} `json:"details,omitempty"`
	IsOperational bool        `json:"isOperational"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: [Code: %s, Status: %d Message:%s Operational:%t] %s", e.Type, e.StatusCode, e.Message, e.IsOperational, e.Details)
}

// GetMessage retrieves the error message from the ErrorType.
func (e *APIError) GetMessage() interface{} {
	return e.Details
}

func NewAPIError(statusCode int, errorType ErrorType, details interface{}, isOperational bool) *APIError {

	return &APIError{
		Type:          errorType.Code,
		StatusCode:    statusCode,
		Message:       errorType.Message,
		Details:       details,
		IsOperational: isOperational,
	}
}

func BadRequestError(errorType ErrorType, details ...interface{}) *APIError {
	detailsVal := interface{}(nil)
	if len(details) > 0 {
		detailsVal = details[0]
	}
	return NewAPIError(http.StatusBadRequest, errorType, detailsVal, true)
}

func AuthenticationError(errorType ErrorType, details ...interface{}) *APIError {
	detailsVal := interface{}(nil)
	if len(details) > 0 {
		detailsVal = details[0]
	}
	return NewAPIError(http.StatusUnauthorized, errorType, detailsVal, true)
}

func NotFoundError(errorType ErrorType, details ...interface{}) *APIError {
	detailsVal := interface{}(nil)
	if len(details) > 0 {
		detailsVal = details[0]
	}
	return NewAPIError(http.StatusNotFound, errorType, detailsVal, true)
}

func InternalServerError(errorType ErrorType, details ...interface{}) *APIError {
	detailsVal := interface{}(nil)
	if len(details) > 0 {
		detailsVal = details[0]
	}
	return NewAPIError(http.StatusInternalServerError, errorType, detailsVal, false)
}
