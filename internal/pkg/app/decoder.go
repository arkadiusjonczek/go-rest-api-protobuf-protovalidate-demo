package app

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/arkadiusjonczek/go-rest-api-protobuf-protovalidate-demo.git/pkg/proto/demo/v1"
)

type Decoder[E Entity] interface {
	Decode(body []byte) (*E, error)
}

var _ Decoder[v1.Customer] = (*CustomerDecoder)(nil)

type CustomerDecoder struct {
	unmarshalOptions *protojson.UnmarshalOptions
	protovalidate    *protovalidate.Validator
}

func NewCustomerDecoder(protovalidate *protovalidate.Validator) *CustomerDecoder {
	return &CustomerDecoder{
		unmarshalOptions: &protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		},
		protovalidate: protovalidate,
	}
}

func (d *CustomerDecoder) Decode(body []byte) (*v1.Customer, error) {
	entry := &v1.Customer{}

	err := d.unmarshalOptions.Unmarshal(body, entry)

	if err != nil {
		return nil, err
	}

	err = d.protovalidate.Validate(entry)
	if err != nil {
		return nil, err
	}

	return entry, nil
}
