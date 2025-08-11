package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	// Init(Config{
	// 	Format:    "text",
	// 	AddSource: true,
	// 	Level:     "debug",
	// })
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")

	// 测试 With 方法
	logger := With("topic", "tunnel")
	logger.Info("info with key")
	logger.Debug("debug with key")
	logger.Warn("warn with key")
	logger.Error("error with key")
}
