package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sashven/sblog/internal/config"
	"github.com/sashven/sblog/internal/content"
)

const shutdownTimeout = 5 * time.Second

type Server struct {
	cfg             config.Config
	router          *gin.Engine
	http            *http.Server
	shutdownTimeout time.Duration
	serve           func() error
	shutdown        func(context.Context) error
}

func NewServer() (*Server, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	posts, err := content.LoadDir(cfg.ContentDir)
	if err != nil {
		return nil, err
	}

	router := newRouter(content.NewIndex(posts))
	httpServer := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}
	return &Server{
		cfg:             cfg,
		router:          router,
		http:            httpServer,
		shutdownTimeout: shutdownTimeout,
		serve:           httpServer.ListenAndServe,
		shutdown:        httpServer.Shutdown,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.serve()
	}()

	select {
	case <-ctx.Done():
		timeout := s.shutdownTimeout
		if timeout == 0 {
			timeout = shutdownTimeout
		}
		shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := s.shutdown(shutdownCtx); err != nil {
			return err
		}
		return ctx.Err()
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

func (s *Server) Close() error {
	return nil
}
