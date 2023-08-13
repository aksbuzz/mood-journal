package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aksbuzz/mood-journal/db"
	"github.com/aksbuzz/mood-journal/server"
	"github.com/aksbuzz/mood-journal/store"
	"github.com/gofiber/fiber/v2/log"
)

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	db := db.New()
	if err := db.Open(ctx); err != nil {
		cancel()
		log.Error(fmt.Printf("Failed to open DB, error: %+v\n", err))
		return
	}

	store := store.New(db.DBInstance)
	s, err := server.New(ctx, store)
	if err != nil {
		cancel()
		log.Error(fmt.Printf("Failed to create new Server, error: %+v\n", err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		log.Info(fmt.Printf("%s received.\n", sig.String()))
		s.Shutdown()
		cancel()
	}()

	if err := s.Start(); err != nil {
		cancel()
		log.Error(fmt.Printf("Failed to start Server, error: %+v\n", err))
		return
	}

	<-ctx.Done()
}
