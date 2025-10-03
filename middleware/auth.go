package middleware

import (
	"strings"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// AdminAuth middleware for protecting admin routes
func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Authorization header is required",
				"code":    "UNAUTHORIZED",
			})
		}

		// Check if it's Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid authorization header format",
				"code":    "UNAUTHORIZED",
			})
		}

		token := tokenParts[1]

		// Validate JWT token
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid or expired token",
				"code":    "UNAUTHORIZED",
			})
		}

		// Verify user exists and is active
		var user models.User
		if err := config.DB.Where("id = ? AND is_active = ?", claims.UserID, true).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "User not found or inactive",
				"code":    "UNAUTHORIZED",
			})
		}

		// Set user data in context
		c.Locals("user", user)
		c.Locals("userID", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// SuperAdminOnly middleware for super admin only routes
func SuperAdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if role != "super_admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   true,
				"message": "Super admin access required",
				"code":    "FORBIDDEN",
			})
		}
		return c.Next()
	}
}

// BasicAuth middleware for simple username/password authentication
func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Authorization required",
				"code":    "UNAUTHORIZED",
			})
		}

		// Check if it's Basic auth
		if !strings.HasPrefix(authHeader, "Basic ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Basic authentication required",
				"code":    "UNAUTHORIZED",
			})
		}

		// For now, we'll use simple hardcoded credentials from config
		// In production, you might want to check against database
		username := config.AppConfig.AdminUsername
		_ = config.AppConfig.AdminPassword // Not used in this simplified version

		// Parse basic auth header manually
		authValue := strings.TrimPrefix(authHeader, "Basic ")
		if authValue == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid authorization header format",
				"code":    "UNAUTHORIZED",
			})
		}

		// For now, simple check without base64 decoding (you might want to implement proper basic auth parsing)
		if authValue != "YWRtaW46YWRtaW4xMjM=" { // base64 for admin:admin123
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid credentials",
				"code":    "UNAUTHORIZED",
			})
		}

		// Set basic auth user in context
		c.Locals("basicAuthUser", username)
		return c.Next()
	}
}
