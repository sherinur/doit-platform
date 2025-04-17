package app

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	logger "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/log"
)

func setupLogger(name string) *logger.Logger {
	loggerProvider := newLoggerProvider()
	logger := loggerProvider.Logger(name, logger.WithInstrumentationVersion("v0.1.0"))

	return &logger
}

func newLoggerProvider() *log.LoggerProvider {
	logExporter, err := stdoutlog.New()
	if err != nil {
		panic(err)
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider
}
