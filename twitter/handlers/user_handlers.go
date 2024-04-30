package handlers

import (
	"net/http"

	"github.com/HsiaoCz/twitter-clone/twitter/storage"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *storage.Store
}

func NewUserHandler(store *storage.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleSignup(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "this is signup",
	})
}

func (u *UserHandler) HandleLogin(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "this is login",
	})
}
func (u *UserHandler) HandleLogout(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "this is logout",
	})
}
