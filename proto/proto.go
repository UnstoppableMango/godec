package proto

import "google.golang.org/protobuf/proto"

var Default = Google

type Message = proto.Message

func Marshal[T Message](v T) ([]byte, error) {
	return Google.Marshal(v)
}

func Unmarshal[T Message](data []byte, v T) error {
	return Google.Unmarshal(data, v)
}
