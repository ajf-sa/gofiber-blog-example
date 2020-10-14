package handlers

import (
	"github.com/alfuhigi/gofiber-blog-example/db"
)

type Handler struct {
	*db.Entity
}

func NewHandler(e *db.Entity) *Handler {
	return &Handler{
		Entity: e,
	}
}
