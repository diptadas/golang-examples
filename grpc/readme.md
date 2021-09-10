### Generate grpc stub and reverse proxy

```
protoc -I /usr/local/include -I . \
-I {$GOPATH}/src/github.com/googleapis/googleapis/ \
--go_out=plugins=grpc:. \
--grpc-gateway_out=logtostderr=true:. \
proto/proto_example.proto
```

### List of processes listening required ports and kill them

```
fuser -k 50051/tcp
fuser -k 8088/tcp
```

### Usage

```
go run server/main.go
go run client/main.go
curl --data "{\"name\": \"dipta\"}" http://localhost:8088/echo
```
