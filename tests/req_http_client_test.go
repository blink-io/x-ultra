package tests

import (
	"fmt"
	"testing"

	"github.com/blink-io/x/internal/testdata"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
)

func TestReq_Get_1(t *testing.T) {
	req.DevMode()                           // 将包名视为 Client 直接调用，启用开发模式
	req.MustGet("https://httpbin.org/uuid") // 将包名视为 Request 直接调用，发起 GET 请求

	req.EnableForceHTTP1() // 强制 HTTP/1.1 看看效果
	req.MustGet("https://httpbin.org/uuid")
}

func TestReq_H3_1(t *testing.T) {
	cc := req.C().
		DevMode().
		SetLogger(zap.S()).
		EnableForceHTTP3().
		SetTLSClientConfig(testdata.GetClientTLSConfig())

	res := cc.R().MustGet("https://localhost:9999/hello")
	body := res.String()

	fmt.Println("body:", body)
}
