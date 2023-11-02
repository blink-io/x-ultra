package flatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterCodec(flatbuffers.FlatbuffersCodec{})
}

type Builder = flatbuffers.Builder
