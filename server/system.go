package server

import "github.com/gofiber/fiber/v2"

func (s *Server) registerSystemRoutes(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Ping received")
	})

	r.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(successResponse(s.Config.App))
	})
}
