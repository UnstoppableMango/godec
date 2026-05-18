package godec

import (
	"github.com/unmango/go/codec"
	"github.com/unstoppablemango/godec/json"
	"github.com/unstoppablemango/godec/proto"
	"github.com/unstoppablemango/godec/yaml"
)

type Codec = codec.Any

var (
	Json  codec.Any = json.Default
	Yaml  codec.Any = yaml.Default
	Proto codec.Any = proto.Default
)
