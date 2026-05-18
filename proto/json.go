package proto

import (
	"io"

	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/internal"
	"google.golang.org/protobuf/encoding/protojson"
)

var Json = NewJson[Message]()

type json[T Message] struct{}

func NewJson[T Message]() json[T] {
	return json[T]{}
}

func (json[T]) Name() string {
	return "google/protojson"
}

func (json[T]) Marshal(v T) ([]byte, error) {
	return protojson.Marshal(v)
}

func (json[T]) Unmarshal(data []byte, v T) error {
	return protojson.Unmarshal(data, v)
}

func (m json[T]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return internal.ReadAll(r, m)
}

func (m json[T]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return internal.WriteAll(w, m)
}
