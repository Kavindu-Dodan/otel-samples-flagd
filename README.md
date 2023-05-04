Multi-language examples for [flagd](https://github.com/open-feature/flagd) simple feature flag evaluations with 
distributed tracing.

### Requirements

- Go 1.18+
- Java 11+
- node for js example
- Docker for docker setup

### How to run

- Start flagd & OpenTelemetry docker setup (see [docker](../docker))
- Run individual language examples (ex:- for go example, `go run main.go`)
- Observe traces in Jaeger(http://127.0.0.1:16686) & metrics in Prometheus(http://localhost:9090)