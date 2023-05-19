module fftest

go 1.20

require (
	github.com/open-feature/go-sdk v1.3.0
	github.com/open-feature/go-sdk-contrib/hooks/open-telemetry v0.2.4
	github.com/open-feature/go-sdk-contrib/providers/flagd v0.1.11
	go.opentelemetry.io/otel v1.15.1
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.38.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.15.1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.15.1
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/sdk/metric v0.38.1
	go.opentelemetry.io/otel/trace v1.15.1
	google.golang.org/grpc v1.54.0
)

require (
	buf.build/gen/go/open-feature/flagd/bufbuild/connect-go v1.5.2-20230222100723-491ee098dd92.1 // indirect
	buf.build/gen/go/open-feature/flagd/protocolbuffers/go v1.29.1-20230317150644-afd1cc2ef580.1 // indirect
	github.com/bufbuild/connect-go v1.7.0 // indirect
	github.com/bufbuild/connect-opentelemetry-go v0.2.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.2 // indirect
	github.com/open-feature/flagd/core v0.5.2 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.15.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.38.0 // indirect
	go.opentelemetry.io/otel/metric v0.38.1 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230221151758-ace64dc21148 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

// Replacement for go-sdk metric hook implementation till release is there
replace github.com/open-feature/go-sdk-contrib/hooks/open-telemetry v0.2.4 =>  /local/dev/go-sdk/path
