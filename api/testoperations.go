package api

import (
	"net/http"
)

// TODO: Remove and replace with actual operation handlers for the api.
// @Summary Test endpoint
// @Description Returns a simple 'Hello, world!' message
// @Success 200 {string} string "Hello, world!"
// @Router /v1/test [get].
func getTestOperation(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
