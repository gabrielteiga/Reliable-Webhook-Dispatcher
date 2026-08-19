package api

import (
	"context"
	"errors"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(address string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    address,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	err := s.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
