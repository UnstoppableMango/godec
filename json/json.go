package json

import (
	"io"

	"github.com/unmango/go/codec"
)

var Default = StdLib

func NewDecoder[T any](r io.Reader) codec.Decoder[T] {
	return NewStdlib[T]().NewDecoder(r)
}

func NewEncoder[T any](w io.Writer) codec.Encoder[T] {
	return NewStdlib[T]().NewEncoder(w)
}

func Marshal[T any](v T) ([]byte, error) {
	return NewStdlib[T]().Marshal(v)
}

func Unmarshal[T any](data []byte, v T) error {
	return NewStdlib[T]().Unmarshal(data, v)
}
