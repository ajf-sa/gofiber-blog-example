package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Index(ctx *fiber.Ctx) error {
	// //TODO remove this
	// _, err := h.CreatePage(ctx.Context(), db.CreatePageParams{Title: "title", Slug: time.Now().String(), Body: "wef"})
	// if err != nil {
	// 	log.Println(err)
	// }

	pages, err := h.ListPage(ctx.Context())

	if err != nil {
		log.Println(err)
	}
	return ctx.Render("index", fiber.Map{"Title": "title", "HomeActive": true, "posts": pages}, "layout")
}

func (h *Handler) GetContact(ctx *fiber.Ctx) error {

	return ctx.Render("contact", fiber.Map{"Title": "title", "ContactActive": true}, "layout")

}

func (h *Handler) PostContact(ctx *fiber.Ctx) error {
	type Contact struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Message string `json:"message"`
	}
	var contact Contact

	err := ctx.BodyParser(&contact)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Render("contact", fiber.Map{"Title": "title", "ContactActive": true, "contact": contact}, "layout")

}

func (h Handler) About(ctx *fiber.Ctx) error {
	return ctx.Render("about", fiber.Map{"Title": "title", "AboutActive": true}, "layout")
}

func (h Handler) Robots(ctx *fiber.Ctx) error {
	return ctx.Render("robots", nil)
}
