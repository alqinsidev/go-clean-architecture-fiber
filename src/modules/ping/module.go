package ping

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/repository"
	transportHttp "alqinsidev.io/go-clean-architecture/src/modules/ping/transport/http"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/usecase"
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	usecase    usecase.UsecaseInterface
	repository *repository.Repository
	router     fiber.Router
}

func Init(app *app.App) *Module {

	repository := repository.Init(app)
	usecase := usecase.Init(app, repository)
	router := transportHttp.Init(app, usecase)

	return &Module{
		repository: repository,
		usecase:    usecase,
		router:     router,
	}
}
