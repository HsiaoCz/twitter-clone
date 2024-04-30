package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleHome(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Welcome to Twitter",
	})
}
