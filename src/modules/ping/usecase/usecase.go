package usecase

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/modules/ping/repository"
)

type UsecaseInterface interface {
	GetHealthCheck() (map[string]string, error)
}

type Usecase struct {
	app        *app.App
	repository *repository.Repository
}

func Init(app *app.App, repository *repository.Repository) *Usecase {
	return &Usecase{
		app:        app,
		repository: repository,
	}
}
