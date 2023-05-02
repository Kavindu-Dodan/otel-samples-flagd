const {OpenFeature} = require('@openfeature/js-sdk');
const {FlagdProvider} = require('@openfeature/flagd-provider');

OpenFeature.setProvider(new FlagdProvider());

// simple flag eval
async function evalFlags() {
    const client = OpenFeature.getClient();

    const decision = await client.getBooleanValue("myBoolFlag", false);
    console.log(decision)
}

evalFlags()

setTimeout(() => {
    console.log('Completed.');
}, 10000)

