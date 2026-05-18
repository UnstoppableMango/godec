package proto

import (
	"fmt"
	"io"

	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/internal"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var Json = json{}

type json struct{}

func (json) Name() string {
	return "google/protojson"
}

func (json) Marshal(v any) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return protojson.Marshal(msg)
	} else {
		return nil, fmt.Errorf("not a proto.Message: %#v", v)
	}
}

func (json) Unmarshal(data []byte, v any) error {
	if msg, ok := v.(proto.Message); ok {
		return protojson.Unmarshal(data, msg)
	} else {
		return fmt.Errorf("not a proto.Message: %#v", v)
	}
}

func (m json) NewDecoder(r io.Reader) codec.Decoder[any] {
	return internal.ReadAll(r, m)
}

func (m json) NewEncoder(w io.Writer) codec.Encoder[any] {
	return internal.WriteAll(w, m)
}
