package internal

import (
	"io"

	"github.com/unmango/go/codec"
)

type castd[T any] struct {
	codec.Decoder[any]
}

func (d castd[T]) Decode(v T) error {
	return d.Decoder.Decode(&v)
}

func AnyD[T any](d codec.Decoder[any]) codec.Decoder[T] {
	return castd[T]{d}
}

type caste[T any] struct {
	codec.Encoder[any]
}

func (e caste[T]) Encode(v T) error {
	return e.Encoder.Encode(v)
}

func AnyE[T any](e codec.Encoder[any]) codec.Encoder[T] {
	return caste[T]{e}
}

type castc[T any, V any] struct {
	codec codec.Codec[V]
}

func (c castc[T, V]) Marshal(v T) ([]byte, error) {
	return c.codec.Marshal(any(v).(V))
}

func (c castc[T, V]) NewDecoder(r io.Reader) codec.Decoder[T] {
	return c.codec.NewDecoder(r).(codec.Decoder[T])
}

func (c castc[T, V]) NewEncoder(w io.Writer) codec.Encoder[T] {
	return c.codec.NewEncoder(w).(codec.Encoder[T])
}

func (c castc[T, V]) Unmarshal(data []byte, v T) error {
	return c.codec.Unmarshal(data, any(v).(V))
}

func Cast[T any, V any, C codec.Codec[V]](c C) codec.Codec[T] {
	return castc[T, V]{c}
}
