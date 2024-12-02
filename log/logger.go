package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New() *Logger {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Reset() {
	l.Logger = slog.Default()
}
