package internal

import (
	"net/http"

	"github.com/ghozilaaa/iso8583-simulator/config"
	"github.com/labstack/echo"
)

type app struct {
	config *config.Config
}

func NewApp(config *config.Config) *app {
	return &app{
		config: config,
	}
}

func (a *app) Start() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	addr := a.config.Server.Host + ":" + a.config.Server.Port
	e.Logger.Fatal(e.Start(addr))
}
