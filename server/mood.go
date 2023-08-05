package server

import (
	"time"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) registerMoodRoutes(r fiber.Router) {
	r.Post("/mood", func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId, err := getUserIdFromContext(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user id", err))
		}
		createMood := &api.CreateMood{}
		if err := c.BodyParser(createMood); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		createMood.UserId = *userId
		createMood.Date = time.Now()
		createMood.Mood = api.Mood(createMood.Mood)
		if err := createMood.Mood.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		mood, err := s.Store.CreateMood(ctx, createMood)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to add mood", err))
		}
		return c.Status(fiber.StatusCreated).JSON(successResponse(mood))
	})

	r.Get("/mood", func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId, err := getUserIdFromContext(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user id", err))
		}
		mood := c.Query("mood")
		timeRange := c.Query("time_range")
		findMood := &api.FindMood{
			UserId:    userId,
			Mood:      api.Mood(mood),
			TimeRange: api.ParseMoodTimeRange(timeRange),
		}
		moods, err := s.Store.ListMoods(ctx, findMood)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to fetch moods", err))
		}
		return c.JSON(successResponse(moods))
	})
}
