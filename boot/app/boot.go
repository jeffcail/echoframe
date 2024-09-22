package app

import (
	"flag"
	"github.com/jeffcail/gtools"
	"github.com/labstack/echo/v4"
)

type App struct {
	c    *echo.Echo
	port string
	cf   string
}

var cf string

func init() {
	flag.StringVar(&cf, "cf", "config", "config file path")
}

func NewApp() *App {
	flag.Parse()
	ap := new(App)
	gm := gtools.Gm

	if ap.c == nil {
		gm.M = gtools.LoadConfig(cf)
	}
	ap.c = echo.New()
	val, ok := gm.M["port"]
	if !ok || val == nil {
		ap.port = ":8090"
	}
	ap.port = val.(string)

	return ap
}

func (a *App) Start() {
	a.c.Logger.Fatal(a.c.Start(a.port))
}
