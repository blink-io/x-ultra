package shared

import (
	"encoding/json"

	"github.com/blink-io/x/kvstore"
)

// Codec KVPair persistence interface.
type Codec interface {
	Encode(kv *kvstore.KVPair) (string, error)
	Decode(b []byte, kv *kvstore.KVPair) error
}

// RawCodec is a simple codec to read and write string.
type RawCodec struct{}

// Encode a KVPair to a string.
func (c RawCodec) Encode(kv *kvstore.KVPair) (string, error) {
	if kv == nil {
		return "", nil
	}

	return string(kv.Value), nil
}

// Decode a byte slice to a KVPair.
func (c RawCodec) Decode(b []byte, kv *kvstore.KVPair) error {
	kv.Value = b

	return nil
}

// JSONCodec is a simple codec to read and write kvstore JSON object.
type JSONCodec struct{}

// Encode a KVPair to a kvstore JSON object.
func (c JSONCodec) Encode(kv *kvstore.KVPair) (string, error) {
	b, err := json.Marshal(kv)
	return string(b), err
}

// Decode a byte slice of kvstore JSON object to a KVPair.
func (c JSONCodec) Decode(b []byte, kv *kvstore.KVPair) error {
	return json.Unmarshal(b, kv)
}
