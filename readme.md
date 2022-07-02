# example.grpc.go.api

## usage

- generate interfaces and documents

```sh
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

protoc \
    --go_out="plugins=grpc:./cat" \
    --doc_out="html,index.html:./cat" \
    ./cat.proto
```

-run

```sh
go run ./
```
