package server

import (
	"context"

	"github.com/aksbuzz/mood-journal/config"
	"github.com/aksbuzz/mood-journal/internal/middleware/requestid"
	"github.com/aksbuzz/mood-journal/store"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	f      *fiber.App
	Store  *store.Store
	Config *config.Config
	Secret string
}

func New(ctx context.Context, store *store.Store, cfg *config.Config) (*Server, error) {
	log.Info("Creating new Server")
	f := fiber.New()
	signingKeySecret := "secret"

	s := &Server{
		f:      f,
		Store:  store,
		Secret: signingKeySecret,
		Config: cfg,
	}

	f.Use(recover.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	f.Use(requestid.New())
	f.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${url} ${status} ${error}\n",
	}))

	apiGroup := f.Group("/api")
	s.registerAuthRoutes(apiGroup)
	s.registerSystemRoutes(apiGroup)

	apiGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(s.Secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse("invalid or expired token", err))
		},
	}))

	s.registerUserRoutes(apiGroup)
	s.registerMoodRoutes(apiGroup)

	return s, nil
}

func (s *Server) Start() error {
	log.Info("Starting Server")
	port := s.Config.Server.Port
	host := s.Config.Server.Host
	return s.f.Listen(host + port)
}

func (s *Server) Shutdown() {
	log.Info("Shutting down!!")
	if err := s.f.Shutdown(); err != nil {
		log.Errorf("Failed to shutdown Server, error: %+v\n", err)
	}

	if err := s.Store.GetDB().Close(); err != nil {
		log.Errorf("Failed to close DB, error: %+v\n", err)
	}

	log.Info("Stopped properly.")
}
