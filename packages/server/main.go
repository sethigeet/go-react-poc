package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sethigeet/go-react-poc/packages/server/database"
	"github.com/sethigeet/go-react-poc/packages/server/util"
)

func main() {
	var err error

	// Load environment variables
	if err := util.LoadEnv(true); err != nil {
		log.Fatalf("there were errors while loading the env file: \n%s", err)
	}

	err = database.Connect(true)
	if err != nil {
		log.Fatalf("there were errors while connecting to the database: \n%s", err)
	}

	srv := InitializeServer()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Attempting to start server on 0.0.0.0:" + DefualtPort)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Panicf("Startup: %s", err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server in a goroutine so that it doesn't block
	go func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Fatalf("Shutdown: %s", err)
		}
	}()

	// Close the database connection
	err = database.Disconnect()
	if err != nil {
		log.Fatalf("Shutdown: error while trying to close the database connection: %s", err)
	}

	log.Println("Shutting down...(Press Ctrl+C again to force)")

	// Listen for another Ctrl+C to force quit
	go func() {
		force := make(chan os.Signal, 1)
		signal.Notify(force, os.Interrupt)
		<-force
		os.Exit(1)
	}()

	// Wait for the server to shutdown properly
	<-ctx.Done()

	os.Exit(0)
}
