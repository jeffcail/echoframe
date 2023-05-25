package _echo

import (
	_uber "github.com/echoframe/pkg/uber"
	"go.uber.org/zap"
)

var EchoLog *EchoLogger

type EchoLogger struct{}

func (*EchoLogger) Write(p []byte) (n int, err error) {
	_uber.EchoScaLog.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
