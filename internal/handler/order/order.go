package handlerorder

import (
	"net/http"

	serviceorder "github.com/echo-scaffolding/internal/service/order"

	"github.com/labstack/echo/v4"
)

//Detail
func Detail(c echo.Context) error {

	detail := serviceorder.Detail()

	return c.JSON(http.StatusOK, detail)
}
