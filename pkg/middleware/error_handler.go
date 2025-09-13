package middleware

import (
	"log"
	"net/http"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/errors"
	"github.com/ctrixcode/ctrix-social-go-backend/pkg/response"
)

// AppHandler is a type that represents an HTTP handler that can return an error.
type AppHandler func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler is a middleware that catches errors returned by handlers
// and panics, then formats them into a standardized JSON error response.
func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				// Log the panic for internal debugging
				log.Printf("PANIC: %v", rvr)

				// Send a generic internal server error response to the client
				apiErr := errors.InternalServerError(errors.ErrInternalServerError)
				response.JSONError(w, apiErr)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// WrapHandler wraps an AppHandler (which returns an error) into a standard http.HandlerFunc.
// It executes the AppHandler and, if an error is returned, calls HandleError.
func WrapHandler(handler AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			HandleError(w, err)
		}
	}
}

// HandleError is a helper function for handlers to return errors.
// The middleware will then catch and process these errors.
// This function is not directly used by the middleware, but by the handlers.
func HandleError(w http.ResponseWriter, err error) {
	apiErr, ok := err.(*errors.APIError)
	if !ok {
		// If it's not a custom APIError, treat it as an internal server error
		apiErr = errors.InternalServerError(errors.ErrInternalServerError, err.Error())
	}

	// Log internal server errors (non-operational errors)
	if !apiErr.IsOperational {
		log.Printf("Internal Server Error: %s (Details: %v)", apiErr.Error(), apiErr.Details)
		// For client, hide specific details for non-operational errors
		apiErr = errors.InternalServerError(errors.ErrInternalServerError) // Generic message
	}

	response.JSONError(w, apiErr)
}
