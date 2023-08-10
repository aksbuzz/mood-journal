package server

import "github.com/gofiber/fiber/v2"

func (s *Server) registerSystemRoutes(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	r.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(successResponse(map[string]interface{}{
			"status": "available",
			"system_info": map[string]string{
				"environment": "development",
				"version":     "1.0.0",
			},
		}))
	})
}
