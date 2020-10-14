package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(ctx *fiber.Ctx) error {
	return ctx.Render("register", fiber.Map{"Title": "title", "RegisterActive": true}, "layout")
}
func (h *Handler) Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{"Title": "title", "LoginActive": true}, "layout")
}

func (h *Handler) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("user")
	return ctx.Redirect("/")

}

func (h *Handler) TryLogin(ctx *fiber.Ctx) error {

	ctx.Cookie(&fiber.Cookie{
		Name:     "user",
		Value:    "randomvalue",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return ctx.Redirect("/")
}
