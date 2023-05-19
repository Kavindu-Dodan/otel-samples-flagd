package main

import (
	"context"
	"fmt"
	"github.com/open-feature/go-sdk/pkg/openfeature"
	api "go.opentelemetry.io/otel/trace"
)

// Flag evaluation logic

func methodA(tracer api.Tracer) {
	client := openfeature.NewClient("methodA")

	spanCtx, span := tracer.Start(context.Background(), "boolEvalLogic")
	defer span.End()

	v, _ := client.BooleanValueDetails(spanCtx, "myBoolFlag", false, openfeature.EvaluationContext{})
	fmt.Printf("flag decision: %t\n", v.Value)
}

func methodB(tracer api.Tracer) {
	client := openfeature.NewClient("methodB")

	spanCtx, span := tracer.Start(context.Background(), "boolEvalLogic")
	defer span.End()

	v, _ := client.BooleanValueDetails(spanCtx, "myBoolFlagB", false, openfeature.EvaluationContext{})
	fmt.Printf("flag decision: %t\n", v.Value)
}

func methodC(tracer api.Tracer) {
	client := openfeature.NewClient("methodC")

	spanCtx, span := tracer.Start(context.Background(), "boolEvalLogic")
	defer span.End()

	v, _ := client.IntValueDetails(spanCtx, "myIntFlag", 1, openfeature.EvaluationContext{})
	fmt.Printf("flag decision: %d\n", v.Value)
}
