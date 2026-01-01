package initialize

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var (
	logger  *slog.Logger
	logFile *os.File
)

// SlogInit 初始化日志系统，同时输出到stdout和日志文件
func SlogInit(level string) {
	// 如果已经初始化过，先关闭之前的日志文件
	if logFile != nil {
		logFile.Sync()
		logFile.Close()
	}

	// 创建logs目录
	logDir := "./logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("创建日志目录失败: %v\n", err)
		return
	}

	// 打开日志文件
	today := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(logDir, fmt.Sprintf("go-cli_%s.log", today))

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开日志文件失败: %v\n", err)
		return
	}

	// 保存文件句柄
	logFile = file

	// 创建MultiWriter，同时写入stdout和文件
	multiWriter := io.MultiWriter(file)

	// 设置日志处理器
	var handler slog.Handler

	// 根据级别设置处理器
	logLevel := slog.LevelInfo
	switch level {
	case "trace", "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	case "fatal", "panic":
		logLevel = slog.LevelError
	}

	handler = slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: logLevel,
	})

	// 创建logger实例
	logger = slog.New(handler)

	// 设置全局logger
	slog.SetDefault(logger)
}

// Close 关闭日志文件
func Close() {
	if logFile != nil {
		logFile.Sync() // 确保数据被写入磁盘
		logFile.Close()
	}
}
