package tests

import (
	"fmt"
	"testing"

	xreq "github.com/blink-io/x/http/client/req"
	"github.com/blink-io/x/http/client/resty"
	"github.com/blink-io/x/internal/testdata"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestReq_Get_1(t *testing.T) {
	cc := xreq.NewClient()
	cc.DevMode()                       // 将包名视为 Client 直接调用，启用开发模式
	cc.Get("https://httpbin.org/uuid") // 将包名视为 Request 直接调用，发起 GET 请求

	cc.EnableForceHTTP1() // 强制 HTTP/1.1 看看效果
	cc.Get("https://httpbin.org/uuid")
}

func TestReq_Head_1(t *testing.T) {
	cc := xreq.NewClient().R().
		MustHead("https://httpbin.org/uuid")
	res := cc.String()
	status := cc.Status

	fmt.Println("status: ", status, ",  res body: ", res)
}

func TestReq_Get_2(t *testing.T) {
	cc := xreq.NewClient().R().
		MustGet("https://httpbin.org/uuid")
	res := cc.String()
	status := cc.Status

	fmt.Println("status: ", status, ",  res body: ", res)
}

func TestReq_H3_1(t *testing.T) {
	cc := xreq.NewClient().
		DevMode().
		SetLogger(zap.S()).
		EnableForceHTTP3().
		SetTLSClientConfig(testdata.GetClientTLSConfig())

	res := cc.R().MustGet("https://localhost:9999/hello")
	body := res.String()

	fmt.Println("body:", body)
}

func TestHTTP3Client_1(t *testing.T) {
	cc := resty.HTTP3Client(testdata.GetClientTLSConfig())
	require.NotNil(t, cc)

	res, err := cc.R().Get("https://localhost:9999/hello")
	require.NoError(t, err)

	fmt.Println("Res body: ", res.String())
}
