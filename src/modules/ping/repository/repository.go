package repository

import "alqinsidev.io/go-clean-architecture/src/app"

type RepositoryInterface interface {
	GetHealthCheck() map[string]string
}

type Repository struct {
	app *app.App
}

func Init(app *app.App) *Repository {
	return &Repository{
		app: app,
	}
}
