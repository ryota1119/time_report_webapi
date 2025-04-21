package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init(env string) {
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	Logger = slog.New(handler)
	log = Logger
	slog.SetDefault(Logger)
}

var log *slog.Logger

func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	log.Warn(fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...any) {
	log.Error(msg, args...)
}

func Fatal(msg string, args ...any) {
	log.Error(msg, args...)
	os.Exit(1)
}

func DebugCtx(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Debug(msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Info(msg, args...)
}

func WarnCtx(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Warn(msg, args...)
}

func ErrorCtx(ctx context.Context, msg string, args ...any) {
	FromContext(ctx).Error(msg, args...)
}
