package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MuxControllers interface {
	ClientController(res http.ResponseWriter, req *http.Request)
}

type FiberControllers interface {
	ClientController(ctx *fiber.Ctx) error
	UpdateConfig(ctx *fiber.Ctx) (err error)
	DeleteConfig(ctx *fiber.Ctx) (err error)
}
