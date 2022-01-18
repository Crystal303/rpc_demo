# proto

```shell
> protoc --go_out=. hello.proto
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
> protoc --go-grpc_out=. hello.proto
```