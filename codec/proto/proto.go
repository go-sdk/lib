package proto

import (
	"google.golang.org/protobuf/proto"
)

type (
	Message = proto.Message

	MarshalOptions   = proto.MarshalOptions
	UnmarshalOptions = proto.UnmarshalOptions
)

var (
	marshalOptions = MarshalOptions{}

	unmarshalOptions = UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func Marshal(m proto.Message) ([]byte, error) {
	return marshalOptions.Marshal(m)
}

func Unmarshal(b []byte, m proto.Message) error {
	return unmarshalOptions.Unmarshal(b, m)
}
