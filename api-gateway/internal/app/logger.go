package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sherinur/doit-platform/api-gateway/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var config zap.Config

	dir := cfg.ZapLogger.Directory
	if dir != "" && !strings.HasSuffix(dir, string(filepath.Separator)) {
		dir += string(filepath.Separator)
	}

	if dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	switch cfg.ZapLogger.Mode {
	case "release":
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{
			"stdout",
			dir + "app.json",
		}
		config.Encoding = "json"
	case "debug":
		config = zap.NewDevelopmentConfig()
		config.OutputPaths = []string{
			"stdout",
			dir + "debug.log",
		}
		config.Encoding = "console"
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "test":
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{
			dir + "test.log",
		}
	default:
		return nil, fmt.Errorf("unknown logging mode: %s", cfg.ZapLogger.Mode)
	}

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
