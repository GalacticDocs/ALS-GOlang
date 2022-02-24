package als_golang

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/andersfylling/disgord"
)

type ILogger struct {
	instance *zap.Logger
}

var _ disgord.Logger = (*ILogger)(nil)

func (log *ILogger) Info(v ...interface{}) {
	log.instance.Info(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *ILogger) Debug(v ...interface{}) {
	log.instance.Debug(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *ILogger) Warn(v ...interface{}) {
	log.instance.Warn(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *ILogger) Error(v ...interface{}) {
	log.instance.Error(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *ILogger) Fatal(v ...interface{}) {
	log.instance.Fatal(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *ILogger) Panic(v ...interface{}) {
	log.instance.Panic(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func InjectableLogger(config zap.Config) *ILogger {
	log, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &ILogger{
		instance: log.With(
			zap.String("lib", "ALS GOlang(v1.0.1)"),
		),
	}
}
