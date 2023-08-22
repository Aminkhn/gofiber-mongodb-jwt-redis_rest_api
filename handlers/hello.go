package handlers

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func HealthChecker(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
