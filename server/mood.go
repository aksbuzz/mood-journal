package server

import (
	"fmt"
	"time"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) registerMoodRoutes(r fiber.Router) {
	r.Post("/mood", func(c *fiber.Ctx) error {
		ctx := c.Context()
		createMood := &api.CreateMood{}
		if err := c.BodyParser(createMood); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("invalid request body, error: %+v", err)})
		}
		createMood.Date = time.Now()
		createMood.Mood = api.Mood(createMood.Mood)
		if err := createMood.Mood.ValidateMood(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("invalid request body, error: %+v", err)})
		}
		mood, err := s.Store.CreateMood(ctx, createMood)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to add mood, error: %+v", err)})
		}
		return c.Status(fiber.StatusCreated).JSON(mood)
	})

	r.Get("/mood", func(c *fiber.Ctx) error {
		ctx := c.Context()
		findMood := &api.FindMood{}
		mood := c.Query("mood")
		date := c.Query("date")
		if mood != "" {
			findMood.Mood = api.Mood(mood)
		}
		findMood.Date = api.ParseMoodDate(date)
		moods, err := s.Store.ListMoods(ctx, findMood)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to fetch moods, error: %+v", err)})
		}
		return c.JSON(fiber.Map{"moods": moods})
	})
}
