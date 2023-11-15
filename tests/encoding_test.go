package tests

import (
	"encoding/base64"
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
		Data: &pb.TestingResponse_Data{
			Action: "testing",
		},
	}
	data, err := proto.Marshal(cpb)
	require.NoError(t, err)
	require.NotNil(t, data)

	//fCode := eproto.FieldNumber(100).Int32(200)
	//fMessage := eproto.FieldNumber(200).String("ok")
	//fDataPayload := eproto.FieldNumber(1).String("testing")
	//fData := eproto.FieldNumber(1).Value(fDataPayload)
	//eproto.MessageRewriter{}
	data2 := testData1()

	fmt.Println("d1: ", base64.StdEncoding.EncodeToString(data))
	fmt.Println("d2: ", base64.StdEncoding.EncodeToString(data2))
}

func TestSegment_Encoding_Proto_2(t *testing.T) {
	data := testData2()
	var mm = new(pb.TestingResponse)
	err := proto.Unmarshal(data, mm)
	require.NoError(t, err)
}

func testData1() []byte {
	md := eproto.AppendVarlen(nil, 1, []byte("testing"))
	m := eproto.AppendVarlen(nil, 1, md)
	m = eproto.AppendVarint(m, 100, 200)
	m = eproto.AppendVarlen(m, 200, []byte("ok"))
	return m
}
func testData2() []byte {
	md := eproto.AppendVarlen(nil, 1, []byte("testing为是一个测试瓦兹"))

	var m eproto.RawMessage
	m = eproto.AppendVarint(m, 100, 200)
	m = eproto.AppendVarlen(m, 200, []byte("ok"))
	m = eproto.AppendVarlen(m, 1, md)
	return m
}
