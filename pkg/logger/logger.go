package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/ctrixcode/ctrix-social-go-backend/pkg/config"
)

// InitLogger initializes the global slog logger based on the provided configuration.
func InitLogger(cfg config.LogConfig) {
	var level slog.Level
	switch strings.ToLower(cfg.Level) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default to info
	}

	var handler slog.Handler
	writer := io.Writer(os.Stdout) // Default output to stdout

	// You could extend this to write to a file or other destinations
	// For example:
	// file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	slog.Error("Failed to open log file", "error", err)
	// } else {
	// 	writer = file
	// }

	switch strings.ToLower(cfg.Format) {
	case "json":
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: level,
		})
	case "text":
		handler = slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: level,
		})
	default:
		handler = slog.NewJSONHandler(writer, &slog.HandlerOptions{ // Default to JSON
			Level: level,
		})
	}

	slog.SetDefault(slog.New(handler))
}
