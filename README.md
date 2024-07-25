# Deepgram CLI

> This project is going through a major redesign effort. If you are looking for our TypeScript CLI which is now deprecated you can find that [here](https://github.com/deepgram-devs/deepgram-cli-legacy).

## Getting an API Key

🔑 To access the Deepgram API you will need a [free Deepgram API Key](https://console.deepgram.com/signup?jump=keys).

## Building the CLI

When you build the Deepgram CLI for the current platform/architecture of your laptop (for example, macOS arm64), you simply need to be at the root of the repo and run:

```bash
go build .
```

> **IMPORTANT:** In order to support multiple platforms, you need to have build mechanisms (ie a Makefile for example) to orchestrate making the binaries for all your target platforms (macOS x86, macOS arm64, Linux amd64, etc, etc).

## External Plugins

### Consuming an External Plugin

You should be able to use the `deepgram-cli` to manage the plugins or download new plugins to your system. You can do this with the `deepgram-cli plugins` command.

```bash
TODO
```

### Manually Installing an External Plugin

For testing or 3rd party purposes, you can also load plugins manually. To do this, copy your Deepgram CLI plugin (this will be a `.so` file extension) into a subfolder called `plugins`.

You will need an accompanying plugin description file, this description file should be the same filename with a `.plugin` extension instead of `.so`.

The contents of the `.plugin` file will look like:

```json
{
    "name": "<your plugin name>",
    "description": "An example plugin description",
    "version": "0.0.1"
}
```

There is an optional field in this JSON that could be used when your root [Cobra CLI Command](https://github.com/spf13/cobra) and initiation function is named something other than `MainCmd` and `InitMain`, respectively. You can specify an optional property named `entrypoint`.

Next time you run the Deepgram CLI, you should see your plugin command available on the command line.

## Development and Contributing

Interested in contributing? We ❤️ pull requests!

To make sure our community is safe for all, be sure to review and agree to our
[Code of Conduct](./CODE_OF_CONDUCT.md). Then see the
[Contribution](./CONTRIBUTING.md) guidelines for more information.

## Getting Help

We love to hear from you so if you have questions, comments or find a bug in the
project, let us know! You can either:

- [Open an issue](https://github.com/deepgram-devs/deepgram-cli/issues) on this repository
- [Join the Deepgram Github Discussions Community](https://github.com/orgs/deepgram/discussions)
- [Join the Deepgram Discord Community](https://discord.gg/xWRaCDBtW4)

## Further Reading

Check out the Developer Documentation at [https://developers.deepgram.com/](https://developers.deepgram.com/)
