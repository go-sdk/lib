package json

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type (
	ProtoMarshalOptions   = protojson.MarshalOptions
	ProtoUnmarshalOptions = protojson.UnmarshalOptions
)

var (
	protoMarshalOptions = ProtoMarshalOptions{
		UseEnumNumbers: true,
	}

	protoUnmarshalOptions = ProtoUnmarshalOptions{
		DiscardUnknown: true,
	}
)

func ProtoMarshal(m proto.Message) ([]byte, error) {
	return protoMarshalOptions.Marshal(m)
}

func ProtoUnmarshal(b []byte, m proto.Message) error {
	return protoUnmarshalOptions.Unmarshal(b, m)
}
