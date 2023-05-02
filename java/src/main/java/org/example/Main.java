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
    public static void main(String[] args) throws InterruptedException {
        //OpenTelemetrySdk telemetrySdk = setupTelemetry("java-manual-telemetry");

        FlagdProvider flagdProvider =
                new FlagdProvider(FlagdOptions.builder().build());

        OpenFeatureAPI api = OpenFeatureAPI.getInstance();
        api.setProvider(flagdProvider);
        Client client = api.getClient();

        for (int i = 0; i < 2; i++) {
            final Boolean boolEval = client.getBooleanValue("myBoolFlag", false);
            System.out.println("eval: " + boolEval);

            Thread.sleep(1000);
        }

        Thread.sleep(5000);
    }


    private static OpenTelemetrySdk setupTelemetry(String appName) {
        Resource resource = Resource.getDefault()
                .merge(Resource.create(Attributes.of(ResourceAttributes.SERVICE_NAME, appName)));

        OtlpGrpcSpanExporter otlpGrpcSpanExporter = OtlpGrpcSpanExporter.builder()
                .setTimeout(Duration.ofSeconds(2)).build();

        SdkTracerProvider sdkTracerProvider = SdkTracerProvider
                .builder()
                .addSpanProcessor(BatchSpanProcessor.builder(otlpGrpcSpanExporter).build())
                .setResource(resource).build();

     return OpenTelemetrySdk.builder()
                .setTracerProvider(sdkTracerProvider)
                .setPropagators(ContextPropagators.create(W3CTraceContextPropagator.getInstance()))
                .buildAndRegisterGlobal();

    }
}

