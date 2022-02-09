package json

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

func MarshalX(v interface{}) ([]byte, error) {
	switch x := v.(type) {
	case json.Marshaler:
		return x.MarshalJSON()
	case proto.Message:
		return ProtoMarshal(x)
	default:
		return Marshal(x)
	}
}

func MustMarshalX(v interface{}) []byte {
	bs, err := MarshalX(v)
	if err != nil {
		panic(err)
	}
	return bs
}

func UnmarshalX(b []byte, v interface{}) error {
	switch x := v.(type) {
	case json.Unmarshaler:
		return x.UnmarshalJSON(b)
	case proto.Message:
		return ProtoUnmarshal(b, x)
	default:
		return Unmarshal(b, v)
	}
}
