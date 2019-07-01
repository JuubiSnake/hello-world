package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/juubisnake/hello-world/pkg/api"
	"github.com/juubisnake/hello-world/pkg/server"
)

func main() {
	// set up new server
	s := server.New()

	// bind our routes
	s.RegisterHandler("/hello-world", api.HelloWorld)

	// listen to traffic and observe any errors thrown
	errChannel := s.ListenAndServe()

	// we'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// we'll shutdown on error or signalled termination
	for {
		select {
		case <-c:
			log.Println("shutting down due to SIGINT, SIGKILL, SIGQUIT or SIGTERM")
			if err := s.Shutdown(); err != nil {
				log.Fatal(err)
			}
			os.Exit(1)
		case err := <-errChannel:
			log.Printf("shutting down due to error: %v\n", err)
			if err := s.Shutdown(); err != nil {
				log.Fatal(err)
			}
			os.Exit(1)
		}
	}
}
