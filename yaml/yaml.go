package yaml

import (
	"io"

	"github.com/unmango/go/codec"
)

var Default codec.Any = Goccy

func NewDecoder(r io.Reader) codec.Decoder[any] {
	return Goccy.NewDecoder(r)
}

func NewEncoder(w io.Writer) codec.Encoder[any] {
	return Goccy.NewEncoder(w)
}

func Marshal(v any) ([]byte, error) {
	return Goccy.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return Goccy.Unmarshal(data, v)
}
