package log

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	time.Time
	activ *slog.Logger
	orig  *slog.Logger
}

func New() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	return &Logger{
		activ: logger,
		orig:  logger,
	}
}

// With sets the given attributes in each output operation. Arguments are converted to attributes as if by Logger.Log.
func (l *Logger) With(args ...any) {
	l.activ = l.orig.With(args...)
}

// ResetAttr remove all set attributes using With
func (l *Logger) ResetAttr() {
	l.activ = l.orig
}

// StartTimer reset receiver' start time
func (l *Logger) StartTimer() {
	l.Time = time.Now()
}

// WithDuration adds duration attr
func (l *Logger) WithDuration() {
	l.With("time", time.Since(l.Time))
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.activ.InfoContext(ctx, msg, args...)
}

func (l *Logger) InfoContextWithTime(ctx context.Context, msg string, args ...any) {
	args = append([]any{"time", time.Since(l.Time)}, args...)
	l.activ.InfoContext(ctx, msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.activ.ErrorContext(ctx, msg, args...)
}

func (l *Logger) ErrorContextWithTime(ctx context.Context, msg string, args ...any) {
	args = append([]any{"time", time.Since(l.Time)}, args...)
	l.activ.ErrorContext(ctx, msg, args...)
}
