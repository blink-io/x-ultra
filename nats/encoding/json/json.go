package json

import (
	"strings"

	"github.com/goccy/go-json"
	"github.com/nats-io/nats.go"
)

const Name = "json"

func init() {
	nats.RegisterEncoder(Name, New())
}

type codec struct {
}

func New() nats.Encoder {
	return &codec{}
}

func (codec) Encode(subject string, v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (codec) Decode(subject string, data []byte, vPtr any) (err error) {
	switch arg := vPtr.(type) {
	case *string:
		// If they want a string and it is a JSON string, strip quotes
		// This allows someone to send a struct but receive as a plain string
		// This cast should be efficient for Go 1.3 and beyond.
		str := string(data)
		if strings.HasPrefix(str, `"`) && strings.HasSuffix(str, `"`) {
			*arg = str[1 : len(str)-1]
		} else {
			*arg = str
		}
	case *[]byte:
		*arg = data
	default:
		err = json.Unmarshal(data, arg)
	}
	return
}

func (codec) Name() string {
	return Name
}
