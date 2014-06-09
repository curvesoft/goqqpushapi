package goqqpushapi

import (
	"crypto/md5"
	"fmt"
	"github.com/curvesoft/beego/httplib"
	"io"
	"net/url"
	"sort"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

const (
	RESTAPI_PUSHSINGLEDEVICE  = "http://openapi.xg.qq.com/v2/push/single_device"
	RESTAPI_PUSHSINGLEACCOUNT = "http://openapi.xg.qq.com/v2/push/single_account"
	RESTAPI_PUSHALLDEVICE     = "http://openapi.xg.qq.com/v2/push/all_device"
	RESTAPI_PUSHTAGS          = "http://openapi.xg.qq.com/v2/push/tags_device"
	RESTAPI_QUERYPUSHSTATUS   = "http://openapi.xg.qq.com/v2/push/get_msg_status"
	RESTAPI_QUERYDEVICECOUNT  = "http://openapi.xg.qq.com/v2/application/get_app_device_num"
	RESTAPI_QUERYTAGS         = "http://openapi.xg.qq.com/v2/tags/query_app_tags"
	RESTAPI_CANCELTIMINGPUSH  = "http://openapi.xg.qq.com/v2/push/cancel_timing_task"

	DEVICE_ALL      = 0
	DEVICE_BROWSER  = 1
	DEVICE_PC       = 2
	DEVICE_ANDROID  = 3
	DEVICE_IOS      = 4
	DEVICE_WINPHONE = 5

	IOSENV_PROD = 1
	IOSENV_DEV  = 2

	TYPE_NOTIFICATION = 1
	TYPE_MESSAGE      = 2
)

type XingePushApp struct {
	accessId  int
	secretKey string
}

func NewXingePushApp(accessId int, secretKey string) *XingePushApp {
	xingeapp := new(XingePushApp)
	xingeapp.accessId = accessId
	xingeapp.secretKey = secretKey
	return xingeapp
}

func (this *XingePushApp) PushSingleDevice(deviceToken string, request Request) Result {
	if request.IsValid() {
		pmap := request.GetParamMap()
		pmap["access_id"] = this.accessId
		pmap["device_token"] = deviceToken

		return this.callRestAPI(RESTAPI_PUSHSINGLEDEVICE, pmap)
	} else {
		return Result{RetCode: -1, RetMsg: "message invalid!"}
	}
}

func (this *XingePushApp) PushSingleAccount(account string, request Request) Result {
	if request.IsValid() {
		pmap := request.GetParamMap()
		pmap["access_id"] = this.accessId
		pmap["account"] = account

		return this.callRestAPI(RESTAPI_PUSHSINGLEACCOUNT, pmap)
	} else {
		return Result{RetCode: -1, RetMsg: "message invalid!"}
	}
}
func (this *XingePushApp) PushAllDevice(request Request) Result {
	if request.IsValid() {
		pmap := request.GetParamMap()
		pmap["access_id"] = this.accessId
		return this.callRestAPI(RESTAPI_PUSHALLDEVICE, pmap)
	} else {
		return Result{RetCode: -1, RetMsg: "message invalid!"}
	}
}

func (this *XingePushApp) callRestAPI(url string, params map[string]interface{}) Result {
	sign := this.makeSign("POST", url, params)
	params["sign"] = sign
	httpreq := httplib.Post(url).Debug(true)
	for k, v := range params {
		httpreq.Param(k, fmt.Sprint(v))
	}
	body, err := httpreq.String()
	println(err)
	println(body)
	return Result{}
}

func (this *XingePushApp) makeSign(method, apiurl string, parammap map[string]interface{}) (sign string) {
	u, err := url.Parse(apiurl)
	if err == nil {
		keys := make([]string, len(parammap))
		i := 0
		for k, _ := range parammap {
			keys[i] = k
			i = i + 1
		}
		sort.Strings(keys)
		paramstr := ""
		sign = ""
		for _, key := range keys {
			paramstr = paramstr + key + "=" + fmt.Sprint(parammap[key])
		}

		sign = Md5(method + u.Host + u.Path + paramstr + this.secretKey)
	}

	return
}
