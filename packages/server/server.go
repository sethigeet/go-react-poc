package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/sethigeet/go-react-poc/packages/server/handler"
	"github.com/sethigeet/go-react-poc/packages/server/middleware"
)

// DefualtPort The default port on which the server should
// run on if the no port is specified in the environment
const DefualtPort = "4000"

// InitializeServer creates an instance of a `http.Server` and `mux.Router`
// with the appropriate route handlers attached to it
func InitializeServer() *http.Server {
	r := mux.NewRouter()

	// middleware
	middleware.Apply(r)

	// Routes
	handler.Apply(r)

	srv := &http.Server{
		Addr: "0.0.0.0:" + DefualtPort,

		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,

		Handler: cors.Default().Handler(r),
	}

	return srv
}
