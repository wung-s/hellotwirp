## Commands

```
// generate the Go files from the proto files
protoc --twirp_out=. --go_out=. rpc/haberdasher/service.proto
protoc --twirp_out=. --go_out=. rpc/helloworld/service.proto
```
