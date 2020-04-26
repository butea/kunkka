package proto

import (
	"github.com/golang/protobuf/proto"
)

type Codec struct{}

func (Codec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (Codec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (Codec) String() string {
	return "proto"
}
