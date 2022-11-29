package router

import (
	"net/http"

	handlerorder "github.com/echo-scaffolding/internal/handler/order"

	confyaml "github.com/echo-scaffolding/conf/yaml"
	"github.com/labstack/echo/v4"
)

//RunHttpServer
func RunHttpServer() {
	e := echo.New()

	orderGroup := e.Group("/v1/order")
	{
		orderGroup.GET("/detail", handlerorder.Detail)
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong...")
	})

	e.Logger.Fatal(e.Start(confyaml.YConf.HTTPBind))
}
