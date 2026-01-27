package proto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type google struct{}

func (google) Name() string {
	return "google/protobuf"
}

func (google) Marshal(v any) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	} else {
		return nil, fmt.Errorf("not a proto.Message: %#v", v)
	}
}

func (google) Unmarshal(data []byte, v any) error {
	if msg, ok := v.(proto.Message); ok {
		return proto.Unmarshal(data, msg)
	} else {
		return fmt.Errorf("not a proto.Message: %#v", v)
	}
}

var Google = google{}
