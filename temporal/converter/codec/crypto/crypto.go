package crypto

import (
	"errors"

	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

type Encryptor interface {
	Encrypt([]byte) ([]byte, error)
}

type Decryptor interface {
	Decrypt([]byte) ([]byte, error)
}

type Cryptor interface {
	Encryptor
	Decryptor
	Name() string
}

const (
	Name = "crypto"

	MetadataEncoding = "binary/crypto+"
)

var _ converter.PayloadCodec = (*codec)(nil)

type Options struct {
	cryptor Cryptor
}

type codec struct {
	ops *Options
}

func New(ops *Options) converter.PayloadCodec {
	return doNew(ops)
}

func doNew(ops *Options) *codec {
	return &codec{ops: ops}
}

func (c *codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	cryptor := c.ops.cryptor
	if cryptor == nil {
		return payloads, errors.New("cryptor is required")
	}

	name := cryptor.Name()
	results := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		encdata, err := cryptor.Encrypt(p.Data)
		if err != nil {
			return payloads, err
		}
		results[i] = &commonpb.Payload{
			Metadata: map[string][]byte{converter.MetadataEncoding: []byte(MetadataEncoding + name)},
			Data:     encdata,
		}
	}
	return results, nil
}

func (c *codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	cryptor := c.ops.cryptor
	if cryptor == nil {
		return payloads, errors.New("cryptor is required")
	}

	name := cryptor.Name()
	results := make([]*commonpb.Payload, len(payloads))
	for i, p := range payloads {
		// Only if it's our encoding
		if string(p.Metadata[MetadataEncoding]) != (MetadataEncoding + name) {
			results[i] = p
			continue
		}
		decdata, err := cryptor.Decrypt(p.Data)
		if err != nil {
			return payloads, err
		}
		results[i] = new(commonpb.Payload)
		err = results[i].Unmarshal(decdata)
		if err != nil {
			return payloads, err
		}
	}
	return results, nil
}

func (c *codec) Name() string {
	return Name
}
