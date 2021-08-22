// Package middleware provides middlewares for the mux router for the API
package middleware

import (
	"github.com/gorilla/mux"
)

func Apply(r *mux.Router) {
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(Logger)
}
