package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sashven/sblog/internal/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("init server: %v", err)
	}
	defer func() {
		if err := server.Close(); err != nil {
			log.Printf("close server: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := server.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("run server: %v", err)
	}
}
