package config

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("http server is shutting down")
	return s.httpServer.Shutdown(ctx)
}
