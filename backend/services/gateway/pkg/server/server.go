package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			Handler:        handler,
			ReadTimeout:    5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (server *Server) Run() {
	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
