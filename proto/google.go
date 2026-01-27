package proto

import (
	"fmt"
	"io"

	"github.com/unmango/go/codec"
	"google.golang.org/protobuf/proto"
)

type google struct{}

func (google) Name() string {
	return "google/protobuf"
}

func (google) NewDecoder(r io.Reader) codec.Decoder[any] {
	panic("google/protobuf does not support streaming decoding")
}

func (google) NewEncoder(w io.Writer) codec.Encoder[any] {
	panic("google/protobuf does not support streaming encoding")
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
