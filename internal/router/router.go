package router

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	handlerauth "github.com/echoframe/internal/handler/auth"

	handleruser "github.com/echoframe/internal/handler/user"

	"github.com/echoframe/conf"

	_middle "github.com/echoframe/common/middle"

	_echo "github.com/echoframe/pkg/echo"

	"go.uber.org/zap"

	uber "github.com/echoframe/pkg/uber"

	"github.com/labstack/echo/v4/middleware"
)

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

	auth := e.Group("/v1/auth")
	{
		auth.POST("/login", handlerauth.AuthLogin)
	}

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
