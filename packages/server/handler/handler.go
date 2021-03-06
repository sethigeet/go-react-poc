// Package handler provides the handlers for the REST API.
// These handlers are the main functions that are called when a request is made
// to the server and they only give the response for the request that was made.
package handler

import "github.com/gorilla/mux"

func Apply(r *mux.Router) {
	// User routes
	(UserHandler{r: r}).Apply()
}
