const apiapi = require('@opentelemetry/api');
const {OpenFeature} = require('@openfeature/js-sdk');
const {FlagdProvider} = require('@openfeature/flagd-provider');
const {api} = require("@opentelemetry/sdk-node");

OpenFeature.setProvider(new FlagdProvider());

// simple flag eval
async function evalFlags() {
    let tracer = api.trace.getTracer("simple-trace");
    const span = tracer.startSpan('client.js:main()');

    const client = OpenFeature.getClient();

    const decision = await client.getBooleanValue("myBoolFlag", false);
    console.log(decision)

    span.end()
}

evalFlags()

setTimeout(() => {
    console.log('Completed.');
}, 10000)

