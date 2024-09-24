package router

import (
	"github.com/jeffcail/echoframe/internal/app/handler"
	"github.com/jeffcail/echoframe/internal/app/middlewares"
	"github.com/jeffcail/echoframe/internal/middles"
	"github.com/labstack/echo/v4"
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func BootApp(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.HEAD, echo.PUT, echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: middlewares.EchoLog,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(middlewares.FrameLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.ConsoleBodyDumpConfig))

	//e.Use(middles.ErrorHandlingMiddleware)
	//e.Use(middles.ResponseMiddleware())
	e.GET("/ping", middles.HandlerMiddleware(handler.ApiDemo))

}
