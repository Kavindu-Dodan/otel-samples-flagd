package main

import (
	"context"
	"fmt"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature" // v1.3.0 from go.mod
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"time"
)

const serviceName = "go-flagd-telemetry"
const collector = "localhost:4317"

func main() {
	// Register OTEL globals
	err := TraceProvider(context.Background(), collector)
	if err != nil {
		fmt.Println("error setting up telemetry" + err.Error())
		return
	}

	// OpenFeature registration - flagd with interceptor for telemetry
	openfeature.SetProvider(flagd.NewProvider(flagd.WithOtelInterceptor(true)))

	client := openfeature.NewClient("app")
	v, _ := client.BooleanValue(
		context.Background(), "myBoolFlag", false, openfeature.EvaluationContext{},
	)

	fmt.Printf("flag decision: %t", v)

	// Enough time to export telemetry
	time.Sleep(10 * time.Second)
}

// TraceProvider register the global trace provider along with propagator for context propagation
func TraceProvider(ctx context.Context, target string) error {
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(target))

	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		return err
	}

	res, err := buildResource(ctx)
	if err != nil {
		return err
	}
	bsp := trace.NewBatchSpanProcessor(traceExp)
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	// set global text propagator for context propagation & global trace provider for traces
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	return nil
}

func buildResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
}
