package errors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Type          ErrorType   `json:"errorType"`
	StatusCode    int         `json:"statusCode"`
	Details       interface{} `json:"details,omitempty"`
	IsOperational bool        `json:"isOperational"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: [Code: %s, Status: %d] %s", e.Type.Code, e.StatusCode, e.GetMessage())
}

func (e *APIError) GetMessage() string {
	if e.Details != nil {
		if msg, ok := e.Details.(string); ok {
			return msg
		}
	}

	return e.Type.Message
}

func NewAPIError(statusCode int, errorType ErrorType, details interface{}, isOperational bool) *APIError {
	return &APIError{
		Type:          errorType,
		StatusCode:    statusCode,
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
