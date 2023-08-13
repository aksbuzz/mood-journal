package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aksbuzz/mood-journal/config"
	"github.com/aksbuzz/mood-journal/db"
	"github.com/aksbuzz/mood-journal/server"
	"github.com/aksbuzz/mood-journal/store"
	"github.com/gofiber/fiber/v2/log"
)

var configFilePath = "config/config.yaml"

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.New(configFilePath)
	if err != nil {
		cancel()
		log.Errorf("Failed to open Config file, error: %+v\n", err)
	}
	db := db.New(cfg)
	if err := db.Open(ctx); err != nil {
		cancel()
		log.Errorf("Failed to open DB, error: %+v\n", err)
		return
	}

	store := store.New(db.DBInstance)
	s, err := server.New(ctx, store, cfg)
	if err != nil {
		cancel()
		log.Errorf("Failed to create new Server, error: %+v\n", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		log.Infof("%s received.\n", sig.String())
		s.Shutdown()
		cancel()
	}()

	if err := s.Start(); err != nil {
		cancel()
		log.Errorf("Failed to start Server, error: %+v\n", err)
		return
	}

	<-ctx.Done()
}
