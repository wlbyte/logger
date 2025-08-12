package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	FormatText = "text"
	FormatJSON = "json"
)

var defaultLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))

type Config struct {
	// 日志级别: "debug", "info", "warn", "error"
	Level string `json:"level" yaml:"level"`
	// 日志格式，支持 "text" 或 "json"
	Format string `json:"format" yaml:"format"`
	// 是否添加源码位置信息
	AddSource bool `json:"addSource" yaml:"addSource"`
	// 日志输出位置，默认为 os.Stdout
	Writer io.Writer
}

func init() {
	Init(Config{
		Format:    "json",
		AddSource: false,
		Level:     "info",
	})
}

func Init(config Config) {
	var level slog.Level
	if os.Getenv("SLOG_LEVEL") != "" {
		config.Level = strings.ToLower(os.Getenv("SLOG_LEVEL"))
	}
	switch config.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handlerOpts := &slog.HandlerOptions{
		Level:     level,
		AddSource: config.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000"))
				}
			}
			if a.Key == slog.SourceKey {
				if source, ok := a.Value.Any().(*slog.Source); ok {
					fileParts := strings.Split(source.File, "/")
					if len(fileParts) > 2 {
						source.File = filepath.Join(fileParts[len(fileParts)-2:]...)
					}
					funcParts := strings.Split(source.Function, ".")
					source.Function = funcParts[len(funcParts)-1]
					a.Value = slog.AnyValue(source)
				}
			}
			return a
		},
	}

	writer := config.Writer
	if writer == nil {
		writer = os.Stdout
	}
	var handler slog.Handler
	switch config.Format {
	case FormatText:
		handler = slog.NewTextHandler(writer, handlerOpts)
	case FormatJSON:
		handler = slog.NewJSONHandler(writer, handlerOpts)
	default:
		handler = slog.NewJSONHandler(writer, handlerOpts)
	}

	defaultLogger = slog.New(handler)
}

func Debug(msg string, args ...any) {
	defaultLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	defaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
}

// DebugContext 等函数支持传入 context
func DebugContext(ctx context.Context, msg string, args ...any) {
	defaultLogger.DebugContext(ctx, msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	defaultLogger.InfoContext(ctx, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	defaultLogger.WarnContext(ctx, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	defaultLogger.ErrorContext(ctx, msg, args...)
}

// With 返回一个新的 logger 实例，该实例包含传入的附加字段。
// 这对于创建带有请求上下文（如 trace_id）的 logger 非常有用。
func With(args ...any) *slog.Logger {
	return defaultLogger.With(args...)
}
