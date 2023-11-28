package pbutil

import (
	"io"

	"github.com/matttproud/golang_protobuf_extensions/v2/pbutil"
	"google.golang.org/protobuf/proto"
)

func ReadDelimited(r io.Reader, m proto.Message) (n int, err error) {
	return pbutil.ReadDelimited(r, m)
}

func WriteDelimited(w io.Writer, m proto.Message) (n int, err error) {
	return pbutil.WriteDelimited(w, m)
}
