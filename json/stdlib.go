package json

import (
	"encoding/json"
	"io"

	"github.com/unmango/go/codec"
)

var StdLib = stdlib{}

type stdlib struct{}

func (stdlib) Name() string {
	return "encoding/json"
}

func (stdlib) NewDecoder(r io.Reader) codec.Decoder[any] {
	return json.NewDecoder(r)
}

func (stdlib) NewEncoder(w io.Writer) codec.Encoder[any] {
	return json.NewEncoder(w)
}

func (stdlib) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (stdlib) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
