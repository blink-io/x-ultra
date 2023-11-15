package tests

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/testdata/pb"
	"github.com/segmentio/encoding/iso8601"
	eproto "github.com/segmentio/encoding/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestSegment_Encoding_1(t *testing.T) {
	str := "2021-10-16T07:55:07+10:00"
	tm, err := iso8601.Parse(str)
	require.NoError(t, err)

	fmt.Println("Time of ISO8601: ", tm)
}

func TestSegment_Encoding_Proto_1(t *testing.T) {
	cpb := &pb.TestingResponse{
		Code:    200,
		Message: "ok",
		//Data: &pb.TestingResponse_Data{
		//	Action: "testing",
		//},
	}
	data, err := proto.Marshal(cpb)
	require.NoError(t, err)
	require.NotNil(t, data)

	//fCode := eproto.FieldNumber(100).Int32(200)
	//fMessage := eproto.FieldNumber(200).String("ok")
	//fDataPayload := eproto.FieldNumber(1).String("testing")
	//fData := eproto.FieldNumber(1).Value(fDataPayload)
	//eproto.MessageRewriter{}

	m := eproto.AppendVarint(nil, 100, 200)
	m = eproto.AppendVarlen(m, 200, []byte("ok"))
	require.NotNil(t, m)

	require.Equal(t, data, m)
}
