package yaml

import (
	"io"

	"github.com/goccy/go-yaml"
	"github.com/unmango/go/codec"
)

var Goccy = goccy{}

type goccy struct{}

func (goccy) Name() string {
	return "goccy/go-yaml"
}

func (goccy) NewDecoder(r io.Reader) codec.Decoder[any] {
	return yaml.NewDecoder(r)
}

func (goccy) NewEncoder(w io.Writer) codec.Encoder[any] {
	return yaml.NewEncoder(w)
}

func (goccy) Marshal(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (goccy) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}
