package main

import (
	"log"
	"net/http"

	"github.com/bufbuild/protovalidate-go"

	"github.com/arkadiusjonczek/go-rest-api-protobuf-protovalidate-demo.git/internal/pkg/app"
	v1 "github.com/arkadiusjonczek/go-rest-api-protobuf-protovalidate-demo.git/pkg/proto/demo/v1"
)

const (
	Addr = ":8080"
)

func main() {
	log.Println("Application started")

	protovalidate, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	decoder := app.NewCustomerDecoder(protovalidate)
	store := app.NewInMemoryStore[v1.Customer]()
	handler := app.NewGenericCrudHandler[v1.Customer](decoder, store)
	router := app.NewCrudRouter[v1.Customer](handler)

	if err := http.ListenAndServe(Addr, router); err != nil {
		log.Fatal(err)
	}

	log.Println("Application stopped")
}
