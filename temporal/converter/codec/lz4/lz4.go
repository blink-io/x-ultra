package lz4

import (
	"bytes"
	"io"

	"github.com/gogo/protobuf/proto"
	"github.com/pierrec/lz4/v4"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

const (
	Name = "lz4"

	MetadataEncoding = "binary/lz4"
)

var _ converter.PayloadCodec = (*codec)(nil)

type codec struct {
	options Options
}

type Options struct {
}

func New(ops Options) converter.PayloadCodec {
	return doNew(ops)
}

func doNew(ops Options) *codec {
	return &codec{options: ops}
}

func (c *codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	results := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		// Marshal and write
		b, err := proto.Marshal(p)
		if err != nil {
			return payloads, err
		}
		var buf bytes.Buffer
		w := lz4.NewWriter(&buf)
		_, err = w.Write(b)
		if closeErr := w.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			return payloads, err
		}
		// Only set if smaller than original amount or has option to always encode
		if buf.Len() < len(b) /* || z.options.AlwaysEncode */ {
			results[i] = &commonpb.Payload{
				Metadata: map[string][]byte{converter.MetadataEncoding: []byte(MetadataEncoding)},
				Data:     buf.Bytes(),
			}
		} else {
			results[i] = p
		}
	}
	return results, nil
}

func (c *codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	results := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		// Only if it's our encoding
		if string(p.Metadata[MetadataEncoding]) != MetadataEncoding {
			results[i] = p
			continue
		}
		r := lz4.NewReader(bytes.NewReader(p.Data))
		// Read all and unmarshal
		b, err := io.ReadAll(r)
		if err != nil {
			return payloads, err
		}
		results[i] = &commonpb.Payload{}
		err = proto.Unmarshal(b, results[i])
		if err != nil {
			return payloads, err
		}
	}
	return results, nil
}

func (c *codec) Name() string {
	return Name
}
