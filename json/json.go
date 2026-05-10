package json

import (
	"io"

	"github.com/unmango/go/codec"
)

var Default codec.Any = StdLib

func NewDecoder(r io.Reader) codec.Decoder[any] {
	return StdLib.NewDecoder(r)
}

func NewEncoder(w io.Writer) codec.Encoder[any] {
	return StdLib.NewEncoder(w)
}

func Marshal(v any) ([]byte, error) {
	return StdLib.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return StdLib.Unmarshal(data, v)
}
