package goqqpushapi

import "time"

type Request struct {
	TimeStamp   int64
	ValidTime   int
	MessageType int
	Message     PushMessage
	ExpireTime  int
	SendTime    time.Time
	MultiPkg    int
	Environment int
}

func NewAndroidRequest(message PushMessage) Request {
	var req Request
	req.ExpireTime = 24*60*60
	req.SendTime = time.Now()
	req.TimeStamp = time.Now().Unix()
	req.ValidTime = 600
	req.MessageType = 1
	req.MultiPkg = 1
	req.Environment = 1
	req.Message = message

	return req
}

func (this *Request) IsValid() bool {
	if this.Message == nil {
		return false
	}
	if !this.Message.IsValid() {
		return false
	}
	if this.MessageType < 0 || this.MessageType > 2 {
		return false
	}
	if this.ExpireTime > 3*24*60*60 {
		return false
	}
	if this.MultiPkg != 0 && this.MultiPkg != 1 {
		return false
	}
	return true
}

func (this *Request) GetParamMap() map[string]interface{} {
	pmap := make(map[string]interface{})
	pmap["timestamp"] = this.TimeStamp
	pmap["valid_time"] = this.ValidTime
	pmap["message_type"] = this.MessageType
	pmap["message"] = this.Message.ToJson()
	pmap["expire_time"] = this.ExpireTime
	if !this.SendTime.IsZero() {
		pmap["send_time"] = this.SendTime.Format("2006-01-02 15:04:05")
	}
	pmap["multi_pkg"] = this.MultiPkg
	pmap["environment"] = this.Environment
	return pmap
}
