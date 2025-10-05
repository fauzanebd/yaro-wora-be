package handlers

import (
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login handles admin login
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Find user by username
	var user models.User
	if err := config.DB.Where("username = ? AND is_active = ?", req.Username, true).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid credentials",
			"code":    "UNAUTHORIZED",
		})
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid credentials",
			"code":    "UNAUTHORIZED",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Username, user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to generate token",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(utils.SimpleAuthResponse{
		Success: true,
		Token:   token,
		Message: "Login successful",
		User: &utils.AuthUser{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		},
	})
}

// Profile returns the current user's profile
func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}
