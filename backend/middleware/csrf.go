package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

// CSRF returns CSRF protection middleware
func CSRF() fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		CookieSecure:   false, // Set to true in production with HTTPS
		CookieHTTPOnly: true,
		Expiration:     1 * 60 * 60, // 1 hour
		KeyGenerator:   csrf.ConfigDefault.KeyGenerator,
		Next: func(c *fiber.Ctx) bool {
			// Skip CSRF for admin endpoints
			path := c.Path()
			return path == "/api/admin/login" ||
			       path == "/api/admin/logout" ||
			       path == "/api/admin/check" ||
			       len(path) >= 14 && path[:14] == "/api/admin/con"
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "CSRF token validation failed",
			})
		},
	})
}
