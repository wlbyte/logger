package main

import (
	"github.com/wlbyte/logger"
)

// logger测试
func main() {
	// logger.Init(logger.Config{
	// 	Format:    "json",
	// 	AddSource: false,
	// 	Level:     "debug",
	// })
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")

	// 测试 With 方法
	logger := logger.With("topic", "tunnel")
	logger.Info("info with key")
	logger.Debug("debug with key")
	logger.Warn("warn with key")
	logger.Error("error with key")
}
