package tests

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/blink-io/x/i18n"
	"github.com/blink-io/x/internal/testdata/pb"
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

func TestSegment_Encoding_3(t *testing.T) {
	pyd1, err1 := os.ReadFile("./testdata/zh-Hans.json")
	require.NoError(t, err1)

	pyd2, err2 := os.ReadFile("./testdata/en-US.json")
	require.NoError(t, err2)

	pyd3 := make([]byte, 0)

	e1 := createEntry("zh-Hans.json", "zh-Hans", true, pyd1)
	e2 := createEntry("en-US.json", "en-US", true, pyd2)
	e3 := createEntry("en-IN.json", "en-IN", false, pyd3)

	//entries := combineEntries(1, e1, e2)
	//require.NotNil(t, entries)

	var m eproto.RawMessage
	m = eproto.AppendVarlen(m, 1, e1)
	m = eproto.AppendVarlen(m, 1, e2)
	m = eproto.AppendVarlen(m, 1, e3)
	m = eproto.AppendVarint(m, 2, 1701148888)

	var res = &i18n.ListLanguagesResponse{}
	require.NotNil(t, res)
	errx := proto.Unmarshal(m, res)
	require.NoError(t, errx)
}

func TestProto_Marshaml(t *testing.T) {
	pyd1, err1 := os.ReadFile("./testdata/zh-Hans.json")
	require.NoError(t, err1)

	pyd2, err2 := os.ReadFile("./testdata/en-US.json")
	require.NoError(t, err2)

	entries := make(map[string]*i18n.LanguageEntry)

	e1 := &i18n.LanguageEntry{
		Path:     "zh-Hans.json",
		Language: "zh-Hans",
		Valid:    true,
		Payload:  pyd1,
	}
	e2 := &i18n.LanguageEntry{
		Path:     "en-US.json",
		Language: "en-US",
		Valid:    true,
		Payload:  pyd2,
	}
	entries["zh-Hans"] = e1
	entries["en-US"] = e2

	var res = &i18n.ListLanguagesResponse{
		Entries:   entries,
		Timestamp: time.Now().Unix(),
	}
	data, errx := proto.Marshal(res)
	require.NoError(t, errx)
	require.NotNil(t, data)

}

func combineEntries(f eproto.FieldNumber, entries ...[]byte) []byte {
	var m eproto.RawMessage
	for _, e := range entries {
		m = eproto.AppendVarlen(m, f, e)
	}
	return m
}

func createEntry(path, language string, valid bool, payload []byte) []byte {
	ed := eproto.AppendVarlen(nil, 1, []byte(path))
	ed = eproto.AppendVarlen(ed, 2, []byte(language))
	ed = eproto.AppendVarint(ed, 3, boolToUint64(valid))
	ed = eproto.AppendVarlen(ed, 20, payload)

	kvd := eproto.AppendVarlen(nil, 1, []byte(language))
	kvd = eproto.AppendVarlen(kvd, 2, ed)
	return kvd
}

func boolToUint64(v bool) uint64 {
	if v {
		return 1
	} else {
		return 0
	}
}
