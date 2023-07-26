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
)

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())
	db := db.NewDB()
	if err := db.Open(); err != nil {
		cancel()
		fmt.Printf("failed to open db, error: %+v\n", err)
		return
	}

	store := store.New(db.DBInstance)
	s, err := server.NewServer(ctx, store)
	if err != nil {
		cancel()
		fmt.Printf("faield to create server, error: %+v\n", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Printf("%s received.\n", sig.String())
		s.Shutdown()
		cancel()
	}()

	if err := s.Start(); err != nil {
		cancel()
		fmt.Printf("failed to start server, error: %+v\n", err)
		return
	}

	<-ctx.Done()
}
