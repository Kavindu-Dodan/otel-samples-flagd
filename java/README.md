Simple Java example to demonstrate distributed traces of flagd. This example use Java manual instrumentation. 
However, it is also possible to setup client application with [automatic instrumentation ](https://opentelemetry.io/docs/instrumentation/java/automatic/)
and receive distributed tracing with flagd.

### How to run ?

1. Run flagd & collector setup through `docker-compose.yaml` located at [docker](../docker) 
2. Build locally with `mvn package`
3. Run the jar with `java -jar ./target/flagd-java-sdk-telemetry-1.0-SNAPSHOT.jar`

### Observe

A new service named `java-flagd-telemetry` will be created and accessible in Jaeger UI - http://127.0.0.1:16686

