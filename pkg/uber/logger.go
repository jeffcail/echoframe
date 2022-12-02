package uber

import (
	"errors"
	"os"
	"sync"

	"github.com/echo-scaffolding/conf"

	"github.com/echo-scaffolding/common/estime"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var EchoScaLog *zap.Logger

// Errors
var (
	RotateFailed  = errors.New("logger Rotate failed err")
	AddFuncFailed = errors.New("adds a func to the Cron err")
)

// Init
func Init() {
	encoder := zapEncoder()
	writeSyncer := zapWriteSyncer()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEndoer := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var Codes []zapcore.Core
	Codes = append(Codes, core)
	Codes = append(Codes, zapcore.NewCore(consoleEndoer, consoleDebug, p))
	c := zapcore.NewTee(Codes...)
	EchoScaLog = zap.New(c, zap.AddCaller())
}

func zapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(estime.LAYOUT)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer() zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:  conf.Config.Logger.Path,
		MaxSize:   conf.Config.Logger.MaxSize,
		MaxAge:    conf.Config.Logger.MaxAge,
		LocalTime: conf.Config.Logger.LocalTime,
		Compress:  conf.Config.Logger.Compress,
	}
	c := cron.New()
	c.AddFunc("0 0 0 1/1 * ?", func() {
		logger.Rotate()
	})

	c.Start()
	return zapcore.AddSync(logger)
}

// NewLogger
func NewLogger() {
	o := &sync.Once{}
	o.Do(Init)
}
