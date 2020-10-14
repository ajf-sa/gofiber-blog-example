package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) Register(ctx *fiber.Ctx) error {
	return ctx.Render("register", fiber.Map{"Title": "title", "RegisterActive": true}, "layout")
}
func (h *Handler) Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{"Title": "title", "LoginActive": true}, "layout")
}

func (h *Handler) Logout(ctx *fiber.Ctx) error {

	return nil

}
