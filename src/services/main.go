package main

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/middlewares"
	"alqinsidev.io/go-clean-architecture/src/modules"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := app.Init()
	router := app.GetFiber()

	router.Use(cors.New(cors.ConfigDefault))
	router.Use(middlewares.RouteLogger(app))

	_ = modules.Init(app)
	_ = app.RunHttp()
}
