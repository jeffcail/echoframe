package handlerauth

import (
	"github.com/echoframe/common/code"
	"github.com/echoframe/internal/input"
	serviceauth "github.com/echoframe/internal/service/auth"
	"github.com/echoframe/utils"
	"github.com/spf13/cast"
)

// AuthLogin
func AuthLogin(c echo.Context) error {
	param := &input.AuthLoginInput{}
	_ = c.Bind(param)
	msg := utils.ValidateParam(param)
	if msg != "" {
		return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
	}
	res, err := serviceauth.AuthLogin(param)
	if err != nil {
		return utils.ToJson(c, utils.Res.Response(false, cast.ToString(err), code.FAILED))
	}

	return utils.ToJson(c, utils.Res.Response(true, "success", code.SUCCESS, res))
}
