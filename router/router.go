package router

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	handleruser "github.com/echo-scaffolding/internal/handler/user"

	"github.com/echo-scaffolding/conf"

	_middle "github.com/echo-scaffolding/common/middle"

	_echo "github.com/echo-scaffolding/pkg/echo"

	"go.uber.org/zap"

	uber "github.com/echo-scaffolding/pkg/uber"

	"github.com/labstack/echo/v4/middleware"

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

	user := e.Group("/v1/user")
	{
		user.POST("/create", handleruser.CreateUser)
		user.GET("/detail/:id", handleruser.UserDetail)
		user.POST("/update", handleruser.UpdateUser)
		user.DELETE("/delete/:id", handleruser.DelUser)
		user.POST("/list", handleruser.UserList)
	}

	e.GET("/ping", func(c echo.Context) error {
		uber.EchoScaLog.Info("Info logger demo")
		uber.EchoScaLog.Info(fmt.Sprintf("Info logger demo :%d", 123))
		uber.EchoScaLog.Error("Error logger demo")
		var err = errors.New("test error demo")
		uber.EchoScaLog.Error(fmt.Sprintf("Error logger demo: %s", "orderno-13546"), zap.Error(err))
		return c.JSON(http.StatusOK, "pong...")
	})

	e.Logger.Fatal(e.Start(conf.Config.HTTPBind))
}
