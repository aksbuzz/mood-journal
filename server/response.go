package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func successResponse(data interface{}) fiber.Map {
	return fiber.Map{
		"status": "success",
		"data":   data,
		"ts":     time.Now().Unix(),
	}
}

func errorResponse(message string, err error) fiber.Map {
	return fiber.Map{
		"status":  "error",
		"message": message,
		"data":    err.Error(),
		"ts":      time.Now().Unix(),
	}
}
