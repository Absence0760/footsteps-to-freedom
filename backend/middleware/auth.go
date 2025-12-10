package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_change_in_production"
	}
	jwtSecret = []byte(secret)
}

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// AdminAuth middleware checks for valid admin authentication
func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from cookie or Authorization header
		token := c.Cookies("admin_token")
		if token == "" {
			authHeader := c.Get("Authorization")
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				token = authHeader[7:]
			}
		}

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: No token provided",
			})
		}

		// Parse and validate token
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !tkn.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token",
			})
		}

		// Check if user is admin
		if !claims.IsAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: Admin access required",
			})
		}

		// Store claims in context for use in handlers
		c.Locals("user", claims)

		return c.Next()
	}
}

// GenerateAdminToken generates a JWT token for admin users
func GenerateAdminToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		IsAdmin:  true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
