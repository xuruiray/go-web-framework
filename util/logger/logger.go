package logger

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xuruiray/go-web-framework/util/config"
	"github.com/xuruiray/go-web-framework/util/request"
	"os"
)

func Init(file string) error {
	var loggerConfig config.LoggerConfig

	err := config.LoadConfig(file, &loggerConfig)
	if err != nil {
		return err
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})

	levelMap := map[string]logrus.Level{
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}

	level, ok := levelMap[loggerConfig.Level]
	if !ok {
		return errors.New(fmt.Sprintf("log level undefined %v", loggerConfig.Level))
	}
	logrus.SetLevel(level)

	switch loggerConfig.Type {
	case "stdio":
		logrus.SetOutput(os.Stdout)
	case "file":
		f, err := os.Create(loggerConfig.Path)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	default:
		return errors.New(fmt.Sprintf("log type undefined %v", loggerConfig.Type))
	}

	return nil
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	logrus.Fatalf(genLog(ctx, format, args...))
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	logrus.Error(genLog(ctx, format, args...))
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	logrus.Info(genLog(ctx, format, args...))
}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	logrus.Warn(genLog(ctx, format, args...))
}
func Debugf(ctx context.Context, format string, args ...interface{}) {
	logrus.Debug(genLog(ctx, format, args...))
}

func genLog(ctx context.Context, format string, args ...interface{}) string {
	lr := request.GetLogRecord(ctx).String()
	return fmt.Sprintf(lr+format, args...)
}
