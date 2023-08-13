package requestid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Config struct {
	ContextKey interface{}
}

var defaultContextKey = "requestid"

func New(config ...Config) func(c *fiber.Ctx) error {
	var cfg Config
	if len(config) < 1 {
		cfg = Config{
			ContextKey: defaultContextKey,
		}
	}
	cfg = config[0]
	if cfg.ContextKey == "" {
		cfg.ContextKey = defaultContextKey
	}

	return func(c *fiber.Ctx) error {
		requestid := uuid.NewString()
		c.Locals(cfg.ContextKey, requestid)
		return c.Next()
	}
}
