Simple JavaScript example to demonstrate distributed traces of flagd

### How to run ?

1. Run flagd & collector setup through `docker-compose.yaml` located at [docker](../docker)
2. Run example with `node --require ./tracer.js client.js`

### Observe

A new service named `js-flagd-telemetry` will be created and accessible in Jaeger UI - http://127.0.0.1:16686

Note - there is an OTEL bug preventing having distributed tracing for backend grpc call

