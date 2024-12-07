package logger

import (
	"context"
	"log/slog"
	"os"
)

var _logger *slog.Logger
var _options *slog.HandlerOptions

type LogLevel slog.Level

const (
	LevelDebug LogLevel = -4
	LevelInfo  LogLevel = 0
	LevelWarn  LogLevel = 4
	LevelError LogLevel = 8
)

type LogValue struct {
	Key   string
	Value any
}

func Var(key string, value any) LogValue {
	return LogValue{Key: key, Value: value}
}

func InitLogger() {
	_logger = slog.New(slog.NewJSONHandler(os.Stdout, _options))
}

func InitLoggerWithLevel(level LogLevel) {
	_options = &slog.HandlerOptions{Level: slog.Level(level)}
}

func convertSlogArgs(args ...LogValue) []slog.Attr {
	slogArgs := make([]slog.Attr, len(args))
	for i, arg := range args {
		slogArgs[i] = slog.Any(arg.Key, arg.Value)
	}
	return slogArgs
}

func logger() *slog.Logger {
	if _logger == nil {
		InitLogger()
	}
	return _logger
}

func Debug(msg string, args ...LogValue) {
	logger().LogAttrs(context.Background(), slog.LevelDebug, msg, convertSlogArgs(args...)...)
}

func Info(msg string, args ...LogValue) {
	logger().LogAttrs(context.Background(), slog.LevelInfo, msg, convertSlogArgs(args...)...)
}

func Warn(msg string, args ...LogValue) {
	logger().LogAttrs(context.Background(), slog.LevelWarn, msg, convertSlogArgs(args...)...)
}

func Error(msg string, args ...LogValue) {
	logger().LogAttrs(context.Background(), slog.LevelError, msg, convertSlogArgs(args...)...)
}
