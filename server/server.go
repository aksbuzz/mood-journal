package server

import (
	"context"
	"fmt"

	"github.com/aksbuzz/mood-journal/store"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	f      *fiber.App
	Store  *store.Store
	Secret string
}

func NewServer(ctx context.Context, store *store.Store) (*Server, error) {
	f := fiber.New()
	signingKeySecret := "secret"

	s := &Server{
		f:      f,
		Store:  store,
		Secret: signingKeySecret,
	}

	f.Use(recover.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	f.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	apiGroup := f.Group("/api")
	s.registerAuthRoutes(apiGroup)

	apiGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(s.Secret)},
	}))

	s.registerUserRoutes(apiGroup)
	s.registerMoodRoutes(apiGroup)
	s.ping()

	return s, nil
}

func (s *Server) Start() error {
	// change on production
	return s.f.Listen("127.0.0.1:8080")
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
