package proto

import (
	"fmt"
	"io"

	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/internal"
	"google.golang.org/protobuf/proto"
)

var Google = NewGoogle[Message]()

type google[T Message] struct{}

func NewGoogle[T Message]() google[T] {
	return google[T]{}
}

func (google[T]) Name() string {
	return "google/protobuf"
}

func (google[T]) Marshal(v T) ([]byte, error) {
	if msg, ok := any(v).(proto.Message); ok {
		return proto.Marshal(msg)
	} else {
		return nil, fmt.Errorf("not a proto.Message: %#v", v)
	}
}

func (google[T]) Unmarshal(data []byte, v T) error {
	return proto.Unmarshal(data, v)
}

func (m google[T]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return internal.ReadAll(r, m)
}

func (m google[T]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return internal.WriteAll(w, m)
}
