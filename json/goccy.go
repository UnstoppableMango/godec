package json

import (
	"io"

	"github.com/goccy/go-json"
	"github.com/unmango/go/codec"
)

var Goccy = goccy{}

type goccy struct{}

func (goccy) Name() string {
	return "goccy/go-json"
}

func (goccy) NewDecoder(r io.Reader) codec.Decoder[any] {
	return json.NewDecoder(r)
}

func (goccy) NewEncoder(w io.Writer) codec.Encoder[any] {
	return json.NewEncoder(w)
}

func (goccy) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (goccy) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
