package app

import (
	"flag"
	"github.com/jeffcail/echoframe/g"
	"github.com/jeffcail/echoframe/vm"
	"github.com/jeffcail/gtools"
	"github.com/labstack/echo/v4"
)

type App struct {
	c    *echo.Echo
	port string
}

var cf string

func init() {
	flag.StringVar(&cf, "cf", "config", "config file path")
}

func NewApp() *App {
	flag.Parse()
	ap := new(App)
	if ap.c == nil {
		g.GM.M = gtools.LoadConfig(cf)
	}
	ap.c = echo.New()

	val := g.GM.Get("port")
	if val.(string) == "" {
		ap.port = ":8090"
	}
	ap.port = val.(string)

	return ap
}

func (a *App) Start() {
	a.do()
	a.c.Logger.Fatal(a.c.Start(a.port))
}

func (a *App) do() {
	vm.NewStore()
}
