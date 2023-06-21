package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"test-go-server/api/router"
	"test-go-server/config"
	"test-go-server/logger"
	"test-go-server/pkg/resource"
)

type Server struct {
	l logger.Logger
	c config.Config
	s *http.Server
}

func New() Server {
	rs := resource.New()
	server := Server{}
	server.l = logger.NewLogger()
	server.c = config.NewConfig()
	r := router.New(server.l, rs)
	server.s = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", server.c.Server.Port),
		Handler: r,
		// ReadTimeout:  c.Server.TimeoutRead,
		// WriteTimeout: c.Server.TimeoutWrite,
		// IdleTimeout:  c.Server.TimeoutIdle,
	}
	return server
}

func (server Server) Run() {
	go func() {
		if err := server.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.l.Error("Server startup failed")
		}
	}()
	server.l.Info("Server is ready to handle requests")
	server.gracefulShutdown()
}

func (server Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	server.l.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.s.SetKeepAlivesEnabled(false)
	if err := server.s.Shutdown(ctx); err != nil {
		server.l.Error("Could not gracefully shutdown the server")
	}
	server.l.Info("Server stopped")
}
