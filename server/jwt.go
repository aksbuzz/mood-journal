package server

import (
	"strconv"
	"time"

	"github.com/aksbuzz/mood-journal/api"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenName  = "access_token"
	issuer     = "moodjournal"
	audience   = "moodjournal.access-token"
	expiration = time.Hour * 24
)

func generateToken(user *api.User, secret []byte) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": issuer,
		"aud": audience,
		"sub": strconv.Itoa(user.ID),
		"exp": jwt.NewNumericDate(time.Now().Add(expiration)),
		"iat": jwt.NewNumericDate(time.Now()),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func getUserIdFromContext(c *fiber.Ctx) (*int, error) {
	sub, err := c.Locals("user").(*jwt.Token).Claims.GetSubject()
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(sub)
	if err != nil {
		return nil, err
	}
	return &userId, nil
}
