Simple Go example to demonstrate distributed traces of flagd

### How to run ?

1. Run flagd & collector setup through `docker-compose.yaml` located at [docker](../docker)
2. Run go application `go run main.go`

### Observe

A new service named `go-flagd-telemetry` will be created and accessible in Jaeger UI - http://127.0.0.1:16686

