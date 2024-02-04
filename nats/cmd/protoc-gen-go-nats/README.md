# Protobuf Generator for NATS Micro Services

This prototype code generator will take protobuf RPC and render services into [NATS](https://nats.io) [micro services](https://pkg.go.dev/github.com/nats-io/nats.go/micro). The reason for the creation of this was to start building a framework on top of the `micro.Service` implementation that is more usable by end users. **STREAM** is not supported in either direction and will be ignored in method signatures.

See the [examples](./examples/) for the proto files input and output created [hello.nats.go](./examples/hello.nats.go).

The code is heavily based on the [Twirp](https://github.com/twitchtv/twirp) implementation. Some improvements may be done internally, and the API may change to allow more `--nats_opt` values.

## Installation

```bash
go install github.com/renevo/protoc-gen-nats@latest
```

## Options

Options in `protoc` can be provided by using the `--nats_opt=<options>` argument. Arguments are provided in a `<key>=<value>` syntax, boolean options do not require a key.

| Key | Value | Description| Example |
|-----|-------|------------|---------|
|`paths`| `import` | Will use the `go_package` path to generate files, this is relative to the working directory (default behavior) | `--nats_opt=paths=import`|
|`paths`| `source_relative` | Will generate the files next to the proto files | `--nats_opt=paths=source_relative`|
|`source_relative`| `true` | Will generate the files next to the proto files | `--nats_opt=source_relative`|

More options will be added in the future.

## Example Usage

Be sure to have the [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) installed.

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/renevo/protoc-gen-nats@latest

protoc --go_out=. --go_opt=paths=source_relative --nats_out=. --nats_opt=source_relative ./examples/hello.proto
```


## Development

In order to actually do development with the plugin, you can use the following:

```bash
go build .
protoc --go_out=. --go_opt=paths=source_relative --nats_out=. --nats_opt=source_relative --plugin=$(pwd)/protoc-gen-nats ./examples/hello.proto
```

This will compile the plugin and then specify it in the protoc command line without having to install it.