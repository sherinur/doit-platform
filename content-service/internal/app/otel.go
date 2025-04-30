package app

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Telemetry struct {
	TracerProvider *sdktrace.TracerProvider
	MetricProvider *metric.MeterProvider
}

func InitTelemetry(ctx context.Context) (*Telemetry, error) {
	traceExp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExp),
	)

	return &Telemetry{
		TracerProvider: tp,
	}, nil
}
