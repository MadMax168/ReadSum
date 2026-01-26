package middleware

import (
	"os"
	"strings"

	"github.com/MadMax168/Readsum/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Authentication
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.NewUnauthorizedError("Missing authorization header")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return utils.NewUnauthorizedError("Server configuration error: JWT secret not set")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.NewUnauthorizedError("Unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return utils.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return utils.NewUnauthorizedError("Invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return utils.NewUnauthorizedError("Invalid user ID in token")
	}

	c.Locals("userID", uint(userID))
	return c.Next()
}
