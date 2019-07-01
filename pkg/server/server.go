package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	s *http.Server
	r *mux.Router
}

// New returns an instance of a server
func New() *server {
	r := mux.NewRouter()
	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	return &server{
		s: s,
		r: r,
	}
}

// RegisterHandler will bind a http handler to our router.
func (s *server) RegisterHandler(p string, h http.HandlerFunc) {
	s.r.HandleFunc(p, h)
}

// ListenAndServe will continuously listen to traffic on all bound routes
// served by a hander.
// Returns an once-only error queue that can be subscribed to, in order to
// halt programs based off errors the server may encounter.
func (s *server) ListenAndServe() <-chan error {
	errChannel := make(chan error, 1)
	go func() {
		if err := s.s.ListenAndServe(); err != nil {
			errChannel <- err
			return
		}
	}()
	return errChannel
}

// Shutdown will gracefully terminate the server after 10 seconds.
func (s *server) Shutdown() error {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := s.s.Shutdown(ctx); err != nil {
		return err
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	return nil
}
