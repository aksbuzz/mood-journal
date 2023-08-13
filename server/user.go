package server

import (
	"fmt"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) registerUserRoutes(r fiber.Router) {
	// Get current user
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

	// Update current user
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
		updateUser := request.UpdateUser(userId)
		user, err := s.Store.UpdateUser(ctx, updateUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to update user", err))
		}
		return c.JSON(successResponse(user))
	})

	// get user settings
	r.Get("/user/settings", func(c *fiber.Ctx) error {
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
		userSettings, err := s.Store.ListUserSettings(ctx, *userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("faied to fetch user settings", err))
		}
		return c.JSON(successResponse(userSettings))
	})

	// upsert user settings
	r.Post("/user/settings", func(c *fiber.Ctx) error {
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
		upsertUserSetting := &api.UpsertUserSetting{}
		if err := c.BodyParser(upsertUserSetting); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		if err := upsertUserSetting.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}
		userSetting, err := s.Store.UpsertUserSetting(ctx, &api.UserSetting{
			UserId:       *userId,
			SettingKey:   upsertUserSetting.SettingKey.String(),
			SettingValue: upsertUserSetting.SettingValue,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to post user settings", err))
		}
		return c.JSON(successResponse(userSetting))
	})
}
