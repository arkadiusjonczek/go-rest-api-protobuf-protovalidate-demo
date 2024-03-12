# go-rest-api-protobuf-protovalidate-demo

This demo is about to demonstrate how to use protobuf and protovalidate with REST API written in go.

## Requirements

To generate the go files for the protobuf definitions you need `buf`.

## Usage

### Protobuf

To generate the go files for the protobuf definitions run:

```shell
buf generate
```

### Server

Start http server with:

```shell
go run cmd/server/main.go
```

### Client

Then request the server with:

```shell
curl localhost:8080/customer
```
