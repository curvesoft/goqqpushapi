package goqqpushapi

import (
	"fmt"
	"testing"
)

var (
	access_id         = 0
	secret_key        = "ukey"
	test_device_token = "utoken"
)

func TestMessageJson(t *testing.T) {
	msg := NewMessageAndroid("title", "content")
	req := NewAndroidRequest(msg)
	xg := NewXingePushApp(access_id, secret_key)
	r := xg.PushAllDevice(req)
	r = xg.PushSingleDevice(test_device_token, req)
	fmt.Printf("result=%v", r)
}

func TestMessageJsonIOS(t *testing.T) {
	msg := NewMessageIOS("titleios")
	req := NewIOSRequest(msg, 2)
	xg := NewXingePushApp(0, "ukey")
	r := xg.PushAllDevice(req)
	fmt.Printf("result=%v", r)
}
