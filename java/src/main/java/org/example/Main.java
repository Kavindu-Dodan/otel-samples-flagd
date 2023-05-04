package org.example;

import dev.openfeature.contrib.providers.flagd.FlagdOptions;
import dev.openfeature.contrib.providers.flagd.FlagdProvider;
import dev.openfeature.sdk.Client;
import dev.openfeature.sdk.OpenFeatureAPI;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.trace.propagation.W3CTraceContextPropagator;
import io.opentelemetry.context.propagation.ContextPropagators;
import io.opentelemetry.exporter.otlp.trace.OtlpGrpcSpanExporter;
import io.opentelemetry.sdk.OpenTelemetrySdk;
import io.opentelemetry.sdk.resources.Resource;
import io.opentelemetry.sdk.trace.SdkTracerProvider;
import io.opentelemetry.sdk.trace.export.BatchSpanProcessor;
import io.opentelemetry.semconv.resource.attributes.ResourceAttributes;

import java.time.Duration;

public class Main {
    private static final String FLAG_KEY = "myBoolFlag";
    private static final String SERVICE = "java-flagd-telemetry";

    public static void main(String[] args) throws InterruptedException {
        final OpenTelemetrySdk telemetrySdk = setupTelemetry(SERVICE);

        // flagd provider with telemetry option
        final FlagdProvider flagdProvider = new FlagdProvider(FlagdOptions.builder().openTelemetry(telemetrySdk).build());

        OpenFeatureAPI api = OpenFeatureAPI.getInstance();
        api.setProvider(flagdProvider);

        final Client client = api.getClient();

        // flag evaluation
        final Boolean boolEval = client.getBooleanValue(FLAG_KEY, false);
        System.out.println("eval: " + boolEval);

        // wait to export everything
        Thread.sleep(5000);
    }


    // OTEL setup
    private static OpenTelemetrySdk setupTelemetry(String appName) {
        Resource resource = Resource.getDefault()
                .merge(Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, appName)));

        OtlpGrpcSpanExporter otlpGrpcSpanExporter = OtlpGrpcSpanExporter.builder()
                .setTimeout(Duration.ofSeconds(2)).build();

        SdkTracerProvider sdkTracerProvider = SdkTracerProvider
                .builder()
                .addSpanProcessor(BatchSpanProcessor.builder(otlpGrpcSpanExporter).build())
                .setResource(resource)
                .build();

     return OpenTelemetrySdk.builder()
                .setTracerProvider(sdkTracerProvider)
                .setPropagators(ContextPropagators.create(W3CTraceContextPropagator.getInstance()))
                .buildAndRegisterGlobal();

    }
}

