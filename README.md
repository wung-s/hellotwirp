A POC using [Twirp](https://github.com/twitchtv/twirp) for RPC (Remote Procudure Call) communication where `serviceA` and `serviceB` talk to each other, with each service acting as
a client as well as a server

## Run the application

This will start both `servicea` and `serviceb` with each calling the other in an infinite loop

```
go run cmd/servicea/main.go
// in another terminal
go run cmd/serviceb/main.go
```

## Pre-requisite for code generation

In the `/rpc` directory, for each package, all the `*.go` files are generated from the `*.proto` file.  
**Important:** The generated `*.go` files must NOT be manually edited. To make changes, update the `*.proto` file and re-run the command below. This will re-generate and update the `*.go` files

Follow the installation steps linked [here](https://twitchtv.github.io/twirp/docs/install.html) before running the code generation commands below

## Command to generate code

```
// generate the Go files from the proto files
protoc --twirp_out=. --go_out=. <path-to-the-proto-file>
```

Example:

```
protoc --twirp_out=. --go_out=. rpc/helloworld/service.proto
```
