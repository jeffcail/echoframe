package handleruser

import (
	"github.com/echoframe/common/code"
	"github.com/echoframe/internal/input"
	serviceuser "github.com/echoframe/internal/service/user"
	"github.com/echoframe/utils"
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
	param := &input.UpdateUserInput{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
	}
	if err := serviceuser.UpdateUser(param); err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}
	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS))
}

// DelUser
func DelUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ToJson(c, utils.Res.Response(false, "参数不完整", code.FAILED))
	}
	d := cast.ToInt64(id)
	if err := serviceuser.DelUser(d); err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}
	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS))
}

// UserList
func UserList(c echo.Context) error {
	param := &input.UserListInput{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
	}
	count, outs, err := serviceuser.UserList(param)
	if err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}
	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS, utils.ResP.Pagination(count, outs)))
}
