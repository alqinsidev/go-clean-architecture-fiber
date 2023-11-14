package http

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/transport/http/handler"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/usecase"
	"github.com/gofiber/fiber/v2"
)

func Init(app *app.App, usecase usecase.UsecaseInterface) fiber.Router {
	router := app.GetFiber().Group("/v1/ping")
	h := handler.InitHandler(app, usecase)

	router.Get("/", h.GetHealthCheck)

	return router
}
