package server

import (
	"context"
	"fmt"

	"github.com/aksbuzz/mood-journal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	f     *fiber.App
	Store *store.Store
}

func NewServer(ctx context.Context, store *store.Store) (*Server, error) {
	f := fiber.New()

	s := &Server{
		f:     f,
		Store: store,
	}

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	apiGroup := f.Group("/api")
	s.registerMoodRoutes(apiGroup)
	s.ping()

	return s, nil
}

func (s *Server) Start() error {
	return s.f.Listen(":8080")
}

func (s *Server) Shutdown() {
	if err := s.f.Shutdown(); err != nil {
		fmt.Printf("failed to shutdown server, error: %+v\n", err)
	}

	if err := s.Store.GetDB().Close(); err != nil {
		fmt.Printf("failed to close databse, error: %+v\n", err)
	}

	fmt.Printf("stopped properly\n")
}

func (s *Server) ping() {
	s.f.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
