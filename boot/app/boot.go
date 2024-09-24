package app

import (
	"flag"
	"github.com/jeffcail/echoframe/g"
	"github.com/jeffcail/echoframe/vm"
	"github.com/jeffcail/gtools"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
	vm.Box.Log.Info("", zap.String("1212212", "1222121"))
	a.c.Logger.Fatal(a.c.Start(a.port))
}

func (a *App) do() {
	vm.NewStore()
}
