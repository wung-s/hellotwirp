A POC using [Twirp](https://github.com/twitchtv/twirp) for RPC (Remote Procudure Call) communication where `serviceA` and `serviceB` talk to each other, thus acting both as a client and a server

## Organization

In the `/rpc` directory, for each package, all the `*.go` files are generated from the `*.proto` file

## Run the application

```

go run cmd/servicea/main.go
// in another terminal
go run cmd/serviceb/main.go
```

## Pre-Requisite:

- Follow the installation steps linked [here](https://twitchtv.github.io/twirp/docs/install.html) before running the code generation commands below

## Command to generate code from the `protobuf` files

```
// generate the Go files from the proto files
protoc --twirp_out=. --go_out=. <path-to-the-proto-file>
```

Example:

```

protoc --twirp_out=. --go_out=. rpc/helloworld/service.proto
```
