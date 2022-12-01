package _middle

import (
	"runtime"
	"strings"

	"github.com/google/uuid"

	"github.com/echo-scaffolding/pkg/uber"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var ReqId = "echo-scaffolding"

// ReqLog
func ReqLog() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10)
					l := runtime.Stack(stack, false)
					uber.EchoScaLog.Error("程序崩溃", zap.String("崩溃日志", string(stack[:l])))
				}
			}()
			if !strings.HasPrefix(c.Path(), "/api/") {
				uber.EchoScaLog.Info("请求开始", zap.Any(c.Request().RequestURI, "网页请求"))
				return handlerFunc(c)
			}
			uid := uuid.New().String()
			c.Set(ReqId, uid)
			err := handlerFunc(c)
			return err
		}
	}
}
