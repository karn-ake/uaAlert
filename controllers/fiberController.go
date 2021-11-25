package controllers

import (
	"log"
	"strings"
	"uaAlert/repository"
	"uaAlert/services"

	"github.com/gofiber/fiber/v2"
)

type fiberController struct {
	repo repository.Repository
	serv services.Services
}

func NewFiberController(repo repository.Repository, serv services.Services) FiberControllers {
	return fiberController{repo, serv}
}

func (c fiberController) ClientController(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	cn := strings.ToUpper(ctx.Params("client"))
	fn, err := c.repo.FindbyClientName(cn)
	if err != nil {
		log.Println(errFindClient)
	}
	post, err := c.serv.CheckStatus(cn, fn.LogFile)
	if err != nil {
		log.Println(errCheckstatus, cn)
	}

	ctx.Status(fiber.StatusOK)
	ctx.JSON(post)

	return nil
}
