// Package handlers contains all the routes for the API
package handlers

import (
	"fmt"
	"net/http"
)

// Handler contains all the routes as methods.
// This makes it easy to spread api keys and secrets between your routes.
// In case you need to add one of those said common parts, you just need to add them to your struct definition.
type Handler struct{}

// NewHandler creates and returns a Handler struct
func NewHandler() *Handler {
	return &Handler{}
}

// Helper function for easily writing response messages
func writeResponse(w http.ResponseWriter, status int, key, value string) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	info := []byte(fmt.Sprintf(`{"%v": "%v"}`, key, value))

	_, err := w.Write(info)
	if err != nil {
		return err
	}

	return nil
}
