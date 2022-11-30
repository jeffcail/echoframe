package router

import (
	"net/http"
	"time"

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

	orderGroup := e.Group("/v1/order")
	{
		orderGroup.GET("/detail", handlerorder.Detail)
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong...")
	})

	e.Logger.Fatal(e.Start(confyaml.YConf.HTTPBind))
}
