package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bots-house/birzzha/cli"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		cancel()
	}()

	cli.Run(ctx)
}
