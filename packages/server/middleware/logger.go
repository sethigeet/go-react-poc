package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%s request on %s", r.Method, r.RequestURI)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
