# gRPC Experiment

* An experiment with gRPC http://www.grpc.io/
* simple token based auth.
* self signed TLS cert generated on the fly.
* useful for learning gRPC https://github.com/kelseyhightower/grpc-hello-service

## Compiling and Running

* Services and messages are defined in `protobuf/...`
* Prerequisites from http://www.grpc.io/docs/quickstart/go.html
* If needed then compile the Go libraries from the protobuf definitions with:

```
protoc --proto_path=protobuf/field/ --go_out=plugins=grpc:field protobuf/field/*
protoc --proto_path=protobuf/data/ --go_out=plugins=grpc:data protobuf/data/*
```

### Server

* Integration tests for the server:

```
cd mtr-api
export $(cat env.list | grep = | xargs); go test -v
```

* Build and run the server:

```
cd mtr-api
export $(cat env.list | grep = | xargs); go build && ./mtr-api
```

* It's possible to handle gRPC and HTTP in the same server see https://coreos.com/blog/gRPC-protobufs-swagger.html


### Client

* Build and run the client.
* Server will log messages.
* Client connection should survive server restarts.
 
```
cd mtr-client
export $(cat env.list | grep = | xargs); go build && ./mtr-client
``` 

### Telemetry and Logging

* Can't do middleware like with http.
* Interceptors have been very recently added which should allow for telemetry https://github.com/grpc/grpc-go/issues/240  
