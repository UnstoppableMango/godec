package typed

import "io"

type Codec[T any] interface {
	Marshaller[T]

	NewDecoder(r io.Reader) Decoder[T]
	NewEncoder(w io.Writer) Encoder[T]
}

type Decoder[T any] interface {
	Decode([]byte) (T, error)
}

type Encoder[T any] interface {
	Encode(T) ([]byte, error)
}

type Marshaller[T any] interface {
	Marshal(T) ([]byte, error)
	Unmarshal([]byte, T) error
}

type decoder[T any] struct{ d Decoder[any] }

func (d decoder[T]) Decode(b []byte) (T, error) {
	var zero T
	v, err := d.d.Decode(b)
	if err != nil {
		return zero, err
	}
	return v.(T), nil
}

type encoder[T any] struct{ e Encoder[any] }

func (e encoder[T]) Encode(v T) ([]byte, error) {
	return e.e.Encode(v)
}

type cast[T any] struct{ c Codec[any] }

func Cast[T any, C Codec[any]](c C) Codec[T] {
	return cast[T]{c}
}

// Marshal implements [Codec].
func (c cast[T]) Marshal(v T) ([]byte, error) {
	return c.c.Marshal(v)
}

// NewDecoder implements [Codec].
func (c cast[T]) NewDecoder(r io.Reader) Decoder[T] {
	return decoder[T]{c.c.NewDecoder(r)}
}

// NewEncoder implements [Codec].
func (c cast[T]) NewEncoder(w io.Writer) Encoder[T] {
	return encoder[T]{c.c.NewEncoder(w)}
}

// Unmarshal implements [Codec].
func (c cast[T]) Unmarshal(b []byte, v T) error {
	return c.c.Unmarshal(b, v)
}
