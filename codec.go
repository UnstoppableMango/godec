package godec

import (
	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/json"
	"github.com/unstoppablemango/godec/proto"
	"github.com/unstoppablemango/godec/yaml"
)

type Codec interface {
	codec.Any
	Name() string
}

var (
	Json  Codec = json.Default
	Yaml  Codec = yaml.Default
	Proto Codec = proto.Default
)
