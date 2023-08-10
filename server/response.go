package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Status string

const (
	Success Status = "success"
	Error   Status = "error"
)

func (s Status) String() string {
	return string(s)
}

func successResponse(data interface{}) fiber.Map {
	return fiber.Map{
		"status":    Success.String(),
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}

func errorResponse(message string, err error) fiber.Map {
	return fiber.Map{
		"status":    Error.String(),
		"message":   message,
		"errors":    err.Error(),
		"timestamp": time.Now().Unix(),
	}
}
