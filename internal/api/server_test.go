package api

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestRunReturnsShutdownError(t *testing.T) {
	shutdownErr := context.DeadlineExceeded
	ctx, cancel := context.WithCancel(context.Background())
	server := &Server{
		shutdownTimeout: time.Second,
		serve: func() error {
			select {}
		},
		shutdown: func(context.Context) error {
			return shutdownErr
		},
	}

	cancel()
	err := server.Run(ctx)

	if !errors.Is(err, shutdownErr) {
		t.Fatalf("Run() error = %v, want %v", err, shutdownErr)
	}
	if errors.Is(err, context.Canceled) {
		t.Fatalf("Run() error = %v, should not include context.Canceled when shutdown fails", err)
	}
}

func TestRunReturnsCancellationWhenShutdownSucceeds(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	server := &Server{
		shutdownTimeout: time.Second,
		serve: func() error {
			select {}
		},
		shutdown: func(context.Context) error {
			return nil
		},
	}

	cancel()
	err := server.Run(ctx)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Run() error = %v, want context.Canceled", err)
	}
}

func TestRunReturnsServeError(t *testing.T) {
	serveErr := errors.New("listen failed")
	server := &Server{
		serve: func() error {
			return serveErr
		},
		shutdown: func(context.Context) error {
			t.Fatal("shutdown should not be called when serving fails first")
			return nil
		},
	}

	err := server.Run(context.Background())

	if !errors.Is(err, serveErr) {
		t.Fatalf("Run() error = %v, want %v", err, serveErr)
	}
}

func TestRunIgnoresErrServerClosed(t *testing.T) {
	server := &Server{
		serve: func() error {
			return http.ErrServerClosed
		},
		shutdown: func(context.Context) error {
			t.Fatal("shutdown should not be called when server is already closed")
			return nil
		},
	}

	err := server.Run(context.Background())

	if err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
}
