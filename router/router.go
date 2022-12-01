package router

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	confini "github.com/echo-scaffolding/conf/ini"

	confnacos "github.com/echo-scaffolding/conf/nacos"

	_middle "github.com/echo-scaffolding/common/middle"

	_echo "github.com/echo-scaffolding/pkg/echo"

	"go.uber.org/zap"

	uber "github.com/echo-scaffolding/pkg/uber"

	"github.com/labstack/echo/v4/middleware"

	handlerorder "github.com/echo-scaffolding/internal/handler/order"

	confyaml "github.com/echo-scaffolding/conf/yaml"
	"github.com/labstack/echo/v4"
)

//RunHttpServer
func RunHttpServer() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.HEAD, echo.PUT, echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: _echo.EchoLog,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	e.Use(_middle.ReqLog())
	e.Use(middleware.BodyDumpWithConfig(_middle.DefaultBodyDumpConfig))

	orderGroup := e.Group("/v1/order")
	{
		orderGroup.GET("/detail", handlerorder.Detail)
	}

	e.GET("/ping", func(c echo.Context) error {
		uber.EchoScaLog.Info("Info logger demo")
		uber.EchoScaLog.Info(fmt.Sprintf("Info logger demo :%d", 123))
		uber.EchoScaLog.Error("Error logger demo")
		var err = errors.New("test error demo")
		uber.EchoScaLog.Error(fmt.Sprintf("Error logger demo: %s", "orderno-13546"), zap.Error(err))
		return c.JSON(http.StatusOK, "pong...")
	})

	if confyaml.YConf != nil {
		e.Logger.Fatal(e.Start(confyaml.YConf.HTTPBind))
	} else if confnacos.NConfig().HTTPBind != "" {
		e.Logger.Fatal(e.Start(confnacos.NConfig().HTTPBind))
	} else if confini.Config().HTTPBind != "" {
		e.Logger.Fatal(e.Start(confini.Config().HTTPBind))
	}

}
