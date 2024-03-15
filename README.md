# go-rest-api-protobuf-protovalidate-demo

This demo is about to demonstrate how to use protobuf and protovalidate with REST API written in go.

## Requirements

To generate the go files for the protobuf definitions you need `buf` installed.

## Usage

### Protobuf

To generate the go files to `pkg/proto/demo/v1` for the protobuf definitions located under `proto/demo/v1` run:

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
curl -XPOST --json '{}' localhost:8080/customer
```

## Routes

The following CRUD routes are available:

```shell
GET    localhost:8080/customer
POST   localhost:8080/customer
GET    localhost:8080/customer/{id}
PUT    localhost:8080/customer/{id}
DELETE localhost:8080/customer/{id}
```
