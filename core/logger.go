package core

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func NewLogger() {
	if Logger == nil {
		LogHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		Logger = slog.New(LogHandler)
	}
}

func LogInfo(msg string, args ...any) {
	Logger.Info(msg, args...)
}

func LogWarn(msg string, args ...any) {
	Logger.Warn(msg, args...)
}

func LogError(msg string, args ...any) {
	Logger.Error(msg, args...)
}
