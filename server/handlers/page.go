package handlers

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func shortRand(length int) string {

	return stringWithCharset(length, charset)

}

func (h *Handler) Index(ctx *fiber.Ctx) error {
	//TODO remove this
	// slug := shortRand(6)
	// _, err := h.CreatePage(ctx.Context(), db.CreatePageParams{Title: "عليك السلام ورحمة الله وبركاتة", Slug: slug, Body: "عليك السلام ورحمة الله وبركاتة"})
	// if err != nil {
	// 	log.Println(err)
	// }

	pages, err := h.ListPage(ctx.Context())

	if err != nil {
		log.Println(err)
	}
	return ctx.Render("index", fiber.Map{"WebsitTitle": "title", "HomeActive": true, "posts": pages}, "layout")
}

func (h Handler) PageBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	page, err := h.GetPageBySlug(ctx.Context(), slug)
	if err != nil {
		log.Println(err)
		return ctx.Redirect("/", fiber.StatusPermanentRedirect)
	}
	return ctx.Render("landpage", fiber.Map{"WebsitTitle": "title", "page": page}, "layout")

}

func (h *Handler) GetContact(ctx *fiber.Ctx) error {

	return ctx.Render("contact", fiber.Map{"WebsitTitle": "title", "ContactActive": true}, "layout")

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
	return ctx.Render("contact", fiber.Map{"WebsitTitle": "title", "ContactActive": true, "contact": contact}, "layout")

}

func (h Handler) About(ctx *fiber.Ctx) error {
	return ctx.Render("about", fiber.Map{"WebsitTitle": "title", "AboutActive": true}, "layout")
}

func (h Handler) Robots(ctx *fiber.Ctx) error {
	return ctx.Render("robots", nil)
}

func (h *Handler) CreateNewPage(ctx *fiber.Ctx) error {
	return ctx.SendString("I am Protected !")
}
