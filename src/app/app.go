package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type App struct {
	ctx    context.Context
	logger *Logger
	fiber  *fiber.App
}

func Init() *App {
	fiber := fiber.New()
	ctx := context.Background()
	log := InitLogger()

	_ = InitConfig()

	return &App{
		ctx:    ctx,
		fiber:  fiber,
		logger: log,
	}
}

func (app *App) GetLogger() *Logger {
	return app.logger
}

func (app *App) GetFiber() *fiber.App {
	return app.fiber
}

func (app *App) RunHttp() error {
	port := viper.GetString("APP_PORT")
	host := fmt.Sprintf(`0.0.0.0:%s`, port)

	return app.fiber.Listen(host)
}
