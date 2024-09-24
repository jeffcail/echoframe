package app

import (
	"flag"
	"github.com/jeffcail/echoframe/internal/app/router"
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
		ap.c = echo.New()
		gtools.LoadConfig(cf)
	}
	val := gtools.Gm.Get("port")
	if val.(string) == "" {
		ap.port = ":8090"
	}
	ap.port = val.(string)

	return ap
}

func (a *App) Start() {
	a.do()
	a.Route()
}

func (a *App) do() {
	vm.BootStore()
}

func (a *App) Route() {
	router.BootApp(a.c)
	a.c.Logger.Fatal(a.c.Start(a.port))
}
