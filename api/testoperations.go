package api

import "net/http"

// TODO: Remove and replace with actual operation handlers for the api.
// @Summary Test endpoint
// @Description Returns a simple 'Hello, world!' message
// @Success 200 {string} string "Hello, world!"
// @Router /v1/test [get]
func getTestOperation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}
