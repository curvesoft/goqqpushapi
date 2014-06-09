package goqqpushapi

import (
	"fmt"
	"testing"
)

var (
	access_id         = 0
	secret_key        = "your_keu"
	test_device_token = "test"
)

func TestMessageJson(t *testing.T) {
	msg := NewMessageAndroid("title", "content")
	req := NewAndroidRequest()
	req.MessageType = 1
	req.Message = msg
	xg := NewXingePushApp(access_id, secret_key)
	r := xg.PushAllDevice(req)
	r = xg.PushSingleDevice(test_device_token, req)
	fmt.Printf("result=%v", r)
}
