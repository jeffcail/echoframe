package _xorm

import (
	_uber "github.com/echoframe/pkg/uber"
	"go.uber.org/zap"
)

var XLog *XLogger

type XLogger struct{}

func (*XLogger) Write(p []byte) (n int, err error) {
	_uber.EchoScaLog.Info("数据库操作", zap.String("数据库", string(p)))
	return len(p), nil
}
