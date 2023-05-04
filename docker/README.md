A self-contained docker setup to up and run flagd & OpenTelemetry collector (Jaeger traces and Prometheus metrics).

This work is adopted from official [OTEL collector example](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/examples/demo)

To start - `docker-compose up`
To clean-up - `docker-compose down`

### Accessing deployments 

Jaeger UI - http://127.0.0.1:16686
Prometheus UI - http://localhost:9090