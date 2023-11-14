package modules

import (
	"alqinsidev.io/go-clean-architecture/src/app"
	"alqinsidev.io/go-clean-architecture/src/modules/ping"
)

type Modules struct {
	Ping *ping.Module
}

func Init(app *app.App) *Modules {
	return &Modules{
		Ping: ping.Init(app),
	}
}
