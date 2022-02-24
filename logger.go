package als_golang

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/andersfylling/disgord"
)

type Logger struct {
	instance *zap.Logger
}

var _ disgord.Logger = (*Logger)(nil)

func (log *Logger) Info(v ...interface{}) {
	log.instance.Info(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *Logger) Debug(v ...interface{}) {
	log.instance.Debug(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *Logger) Warn(v ...interface{}) {
	log.instance.Warn(fmt.Sprint(v...))
	_ = log.instance.Sync()
}

func (log *Logger) Error(v ...interface{}) {
	log.instance.Error(fmt.Sprint(v...))
	_ = log.instance.Sync()
}
