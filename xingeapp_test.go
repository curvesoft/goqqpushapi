package goqqpushapi

import (
	"fmt"
	"testing"
)

var (
	access_id         = 2100029665
	secret_key        = "717489ce7e2d1b37bc09799f960a6924"
	test_device_token = "d5a8b9480c89dcbb4c38b7b5c18b9ae1b5b8ee5d"
)

func TestMessageJson(t *testing.T) {
	msg := NewMessageAndroid("title", "content")
	req := NewAndroidRequest(msg)
	xg := NewXingePushApp(access_id, secret_key)
	r := xg.PushAllDevice(req)
	r = xg.PushSingleDevice(test_device_token, req)
	fmt.Printf("result=%v", r)
}
