package internal

import (
	"io"

	"github.com/unmango/go/codec"
)

type readerAll[T any] struct {
	r io.Reader
	m codec.Marshaler[T]
}

func (d *readerAll[T]) Decode(v T) error {
	data, err := io.ReadAll(d.r)
	if err != nil {
		return err
	}
	return d.m.Unmarshal(data, v)
}

// ReadAll returns a Decoder that reads the entire input into memory and then unmarshals it using the provided Marshaler.
// It is intended to be used to shim codecs that do not support streaming.
func ReadAll[T any, M codec.Marshaler[T]](r io.Reader, m M) codec.Decoder[T] {
	return &readerAll[T]{r, m}
}

type writerAll[T any] struct {
	w io.Writer
	m codec.Marshaler[T]
}

func (e *writerAll[T]) Encode(v T) error {
	data, err := e.m.Marshal(v)
	if err != nil {
		return err
	}
	_, err = e.w.Write(data)
	return err
}

// WriteAll returns an Encoder that marshals the value into memory and then writes it to the provided Writer.
// It is intended to be used to shim codecs that do not support streaming.
func WriteAll[T any, M codec.Marshaler[T]](w io.Writer, m M) codec.Encoder[T] {
	return &writerAll[T]{w, m}
}
