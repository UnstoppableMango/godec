package json

import (
	"encoding/json"
	"io"

	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/internal"
)

var StdLib = NewStdlib[any]()

type stdlib[T any] struct{}

func NewStdlib[T any]() stdlib[T] {
	return stdlib[T]{}
}

func (stdlib[T]) Name() string {
	return "encoding/json"
}

func (stdlib[T]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return internal.AnyD[T](json.NewDecoder(r))
}

func (stdlib[T]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return internal.AnyE[T](json.NewEncoder(w))
}

func (stdlib[T]) Marshal(v T) ([]byte, error) {
	return json.Marshal(v)
}

func (stdlib[T]) Unmarshal(data []byte, v T) error {
	return json.Unmarshal(data, v)
}
