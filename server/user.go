package server

import (
	"fmt"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) registerUserRoutes(r fiber.Router) {
	r.Get("/user", func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId, err := getUserIdFromContext(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user id", err))
		}
		user, err := s.Store.FindUser(ctx, &api.FindUser{
			ID: userId,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user", err))
		}
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse("failed to get user", fmt.Errorf("user does not exists in the database")))
		}
		return c.JSON(successResponse(user))
	})

	r.Patch("/user", func(c *fiber.Ctx) error {
		ctx := c.Context()
		userId, err := getUserIdFromContext(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user id", err))
		}
		currUser, err := s.Store.FindUser(ctx, &api.FindUser{
			ID: userId,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to get user", err))
		}
		if currUser == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse("failed to get user", fmt.Errorf("user does not exists in the database")))
		}
		request := &api.PatchUserRequest{}
		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		if err := request.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		updateUser := request.GetUpdateUser(userId)
		user, err := s.Store.UpdateUser(ctx, updateUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to update user", err))
		}
		return c.JSON(successResponse(user))
	})
}
