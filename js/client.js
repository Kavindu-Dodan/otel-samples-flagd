const apiapi = require('@opentelemetry/api');
const {OpenFeature} = require('@openfeature/js-sdk');
const {FlagdProvider} = require('@openfeature/flagd-provider');
const {api} = require("@opentelemetry/sdk-node");

const flagkey = "myBoolFlag"

OpenFeature.setProvider(new FlagdProvider());

// simple flag eval
async function evalFlags() {
    const tracer = api.trace.getTracer("flag-eval");

    const span = tracer.startSpan('eval-flag');

    const client = OpenFeature.getClient();
    const decision = await client.getBooleanValue(flagkey, false);
    console.log("flag decision: " + decision)

    span.end()
}

evalFlags()
// wait for exporting to complete
setTimeout(() => {
    console.log('complete');
}, 10000)

