package server

import (
	"fmt"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) registerAuthRoutes(r fiber.Router) {
	r.Post("/auth/signin", func(c *fiber.Ctx) error {
		ctx := c.Context()
		signin := &api.SignIn{}
		if err := c.BodyParser(signin); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}

		user, err := s.Store.FindUser(ctx, &api.FindUser{
			UserName: &signin.Username,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("incorrect login credentials", err))
		}
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse("incorrect login credentials", fmt.Errorf("user does not exists in the database")))
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(signin.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse("incorrect login credentials", err))
		}

		t, err := generateToken(user, []byte(s.Secret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to generate access token", err))
		}

		return c.JSON(successResponse(map[string]interface{}{
			"user":  user,
			"token": t,
		}))
	})

	r.Post("/auth/signup", func(c *fiber.Ctx) error {
		ctx := c.Context()
		signup := &api.SignUp{}
		if err := c.BodyParser(signup); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errorResponse("invalid request body", err))
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to generate password hash", err))
		}

		createUser := &api.User{
			Username:     signup.Username,
			PasswordHash: string(passwordHash),
		}

		user, err := s.Store.CreateUser(ctx, createUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to create user", err))
		}

		t, err := generateToken(user, []byte(s.Secret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(errorResponse("failed to generate access token", err))
		}

		return c.JSON(successResponse(map[string]interface{}{
			"user":  user,
			"token": t,
		}))
	})
}
