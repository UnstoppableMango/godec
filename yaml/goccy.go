package yaml

import (
	"io"

	"github.com/goccy/go-yaml"
	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/internal"
)

var Goccy = NewGoccy[any]()

type goccy[T any] struct{}

func NewGoccy[T any]() goccy[T] {
	return goccy[T]{}
}

func (goccy[T]) Name() string {
	return "goccy/go-yaml"
}

func (goccy[T]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return internal.AnyD[T](yaml.NewDecoder(r))
}

func (goccy[T]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return internal.AnyE[T](yaml.NewEncoder(w))
}

func (goccy[T]) Marshal(v T) ([]byte, error) {
	return yaml.Marshal(v)
}

func (goccy[T]) Unmarshal(data []byte, v T) error {
	return yaml.Unmarshal(data, v)
}
