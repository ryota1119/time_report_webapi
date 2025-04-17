package logger

import (
	"context"
	"log/slog"
)

type contextKey struct{}

var ctxKey = &contextKey{}

func WithContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, l)
}

func FromContext(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(ctxKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return l
}
