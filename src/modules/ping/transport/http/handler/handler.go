package handler

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/usecase"
)

type Handler struct {
	app     *app.App
	usecase usecase.UsecaseInterface
}

var (
	MethodGetHealthCheck = `get-health-check`
)

func InitHandler(app *app.App, usecase usecase.UsecaseInterface) *Handler {
	return &Handler{
		app:     app,
		usecase: usecase,
	}
}
