package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

func Error(ctx context.Context, format string, args ...interface{}) {
	logrus.Error()
}
func Info(ctx context.Context, format string, args ...interface{})  {}
func Warn(ctx context.Context, format string, args ...interface{})  {}
func Debug(ctx context.Context, format string, args ...interface{}) {}
