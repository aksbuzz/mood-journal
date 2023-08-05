package server

import (
	"strconv"
	"time"

	"github.com/aksbuzz/mood-journal/api"
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

// func getUserFromContext(c *fiber.Ctx) (int, error) {
// 	user := c.Locals("user").(*jwt.Token)
// 	sub, err := user.Claims.GetSubject()
// 	if err != nil {
// 		return 0, err
// 	}
// }
