package proto

import (
	"github.com/unmango/go/codec"
)

var Default codec.Marshaler[any] = Google

func Marshal(v any) ([]byte, error) {
	return Google.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return Google.Unmarshal(data, v)
}
