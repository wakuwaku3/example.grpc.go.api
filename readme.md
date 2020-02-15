# example.grpc.go.api

## usage

- generate interfaces and documents

```sh
protoc \
    --go_out="plugins=grpc:./cat" \
    --doc_out="html,index.html:./cat" \
    ./cat.proto
```

-run

```sh
go run ./
```
