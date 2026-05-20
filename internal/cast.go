package internal

import (
	"io"

	"github.com/unmango/go/codec"
)

type Decoder[T any] struct {
	codec.Decoder[any]
}

func (d Decoder[T]) Decode(v T) error {
	return d.Decoder.Decode(&v)
}

type Encoder[T any] struct {
	codec.Encoder[any]
}

func (e Encoder[T]) Encode(v T) error {
	return e.Encoder.Encode(v)
}

type cast[T any, V any] struct {
	codec codec.Codec[V]
}

func (c cast[T, V]) Marshal(v T) ([]byte, error) {
	return c.codec.Marshal(any(v).(V))
}

func (c cast[T, V]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return c.codec.NewDecoder(r).(codec.Decoder[T])
}

func (c cast[T, V]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return c.codec.NewEncoder(w).(codec.Encoder[T])
}

func (c cast[T, V]) Unmarshal(data []byte, v T) error {
	return c.codec.Unmarshal(data, any(v).(V))
}

func Cast[T any, V any, C codec.Codec[V]](c C) codec.Codec[T] {
	return cast[T, V]{c}
}
