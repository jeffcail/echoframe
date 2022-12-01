package _middle

import (
	"strings"

	"github.com/echo-scaffolding/pkg/uber"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var DefaultBodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: BodyDumpSkipper,
	Handler: func(c echo.Context, bytes []byte, bytes2 []byte) {
		if !strings.HasPrefix(c.Path(), "/api/") {
			return
		}
		uid := c.Get(ReqId).(string)
		uber.EchoScaLog.Info("请求结束", zap.String("请求UID", uid), zap.String(c.Request().RequestURI, string(bytes2)))
	},
}

func BodyDumpSkipper(c echo.Context) bool {
	if !strings.HasPrefix(c.Path(), "/api/") {
		return true
	}
	return false
}
