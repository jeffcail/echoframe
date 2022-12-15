package handleruser

import (
	"github.com/echo-scaffolding/common/code"
	"github.com/echo-scaffolding/internal/handler/input"
	"github.com/echo-scaffolding/utils"
	"github.com/labstack/echo/v4"
)

// CreateUser
func CreateUser(c echo.Context) error {
	param := &input.CreateUserInput{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return utils.ToXml(c, utils.Res.Response(false, msg, code.FAILED))
	}
	return nil
}

// UserDetail
func UserDetail(c echo.Context) error {

	return nil
}

// UpdateUser
func UpdateUser(c echo.Context) error {

	return nil
}

// DelUser
func DelUser(c echo.Context) error {

	return nil
}

// UserList
func UserList(c echo.Context) error {

	return nil
}
