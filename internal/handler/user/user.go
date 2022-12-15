package handleruser

import (
	"github.com/echo-scaffolding/common/code"
	"github.com/echo-scaffolding/internal/input"
	serviceuser "github.com/echo-scaffolding/internal/service/user"
	"github.com/echo-scaffolding/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// CreateUser
func CreateUser(c echo.Context) error {
	param := &input.CreateUserInput{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
	}
	if err := serviceuser.CreateUser(param); err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}
	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS))
}

// UserDetail
func UserDetail(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ToJson(c, utils.Res.Response(false, "参数不完整", code.FAILED))
	}
	d := cast.ToInt64(id)
	u, err := serviceuser.UserDetail(d)
	if err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}
	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS, u))
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
