package app

import (
	"fmt"

	"github.com/sherinur/doit-platform/quiz-service/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var config zap.Config

	switch cfg.ZapLogger.Mode {
	case "release":
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{
			"stdout",
			cfg.ZapLogger.Directory + "app.log",
		}
	case "debug":
		config = zap.NewDevelopmentConfig()
		config.OutputPaths = []string{
			"stdout",
			cfg.ZapLogger.Directory + "debug.log",
		}
	case "test":
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{
			cfg.ZapLogger.Directory + "test.log",
		}
	default:
		return nil, fmt.Errorf("unknown logging mode: %s", cfg.ZapLogger.Mode)
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
