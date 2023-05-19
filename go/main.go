package main

import (
	"context"
	"fmt"
	hooks "github.com/open-feature/go-sdk-contrib/hooks/open-telemetry/pkg"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
	"go.opentelemetry.io/otel"
	"time"
)

const (
	serviceName     = "go-flagd-telemetry"
	collectorTarget = "localhost:4317"
)

// initializer
func init() {
	ctx := context.Background()

	// Register OTEL globals
	err := TraceProvider(ctx, collectorTarget)
	if err != nil {
		fmt.Println("error setting up telemetry" + err.Error())
		return
	}

	// Derive metric exporter
	reader, err := MetricReader(ctx, collectorTarget, 2*time.Second)
	if err != nil {
		fmt.Println("error setting up metrics" + err.Error())
		return
	}

	// Derive metric hook from reader
	metricsHook, err := hooks.NewMetricsHook(reader)
	if err != nil {
		fmt.Println("error setting up metrics" + err.Error())
		return
	}

	// Register OpenFeature API hooks
	openfeature.AddHooks(hooks.NewHook())
	openfeature.AddHooks(metricsHook)

	// Register provider - flagd with interceptor for telemetry
	openfeature.SetProvider(flagd.NewProvider(flagd.WithOtelInterceptor(true)))

}

func main() {
	// go eval tracer
	tracer := otel.Tracer("go-eval")

	for true {
		go methodA(tracer)
		go methodB(tracer)
		go methodC(tracer)

		time.Sleep(500 * time.Millisecond)
	}
}
