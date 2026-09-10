# Godec

Codec libraries as values.

## Usage

```go
import (
    "github/unstoppablemango/godec"
    "github/unstoppablemango/godec/json"
    "github/unstoppablemango/godec/proto"
    "github/unstoppablemango/godec/yaml"
)

var codec godec.Codec

codec = godec.Json
codec = godec.Yaml
codec = godec.Proto

codec = json.StdLib
codec = yaml.Goccy
codec = proto.Google
```

`godec` does as little as possible, for example:

```go
import "github/unstoppablemango/godec/json"

_ = json.StdLib.NewDecoder(r)
```

is equivalent to:

```go
import "encoding/json"

_ = json.NewDecoder(r)
```

Errors will be returned where necessary to satisfy interface requirements.
For example in the `proto` package:

```go
import "google.golang.org/protobuf/proto"

// ...

func (google) Marshal(v any) ([]byte, error) {
	if msg, ok := v.(proto.Message); ok {
		return proto.Marshal(msg)
	} else {
		return nil, fmt.Errorf("not a proto.Message: %#v", v)
	}
}
```
