package handlers

import (
	"os"
	"time"

	"app/middleware"

	"github.com/gofiber/fiber/v2"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminLogin handles admin authentication
func AdminLogin(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get credentials from environment
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	// Default credentials for development (CHANGE IN PRODUCTION!)
	if adminUsername == "" {
		adminUsername = "admin"
	}
	if adminPassword == "" {
		adminPassword = "admin123"
	}

	// Verify credentials
	if req.Username != adminUsername || req.Password != adminPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token, err := middleware.GenerateAdminToken(req.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set token in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "admin_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: "Strict",
	})

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// AdminLogout handles admin logout
func AdminLogout(c *fiber.Ctx) error {
	// Clear the admin token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "admin_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}

// CheckAuth verifies if the user is authenticated
func CheckAuth(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"authenticated": false,
		})
	}

	return c.JSON(fiber.Map{
		"authenticated": true,
		"user":          user,
	})
}
