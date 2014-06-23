package goqqpushapi

import "encoding/json"

type PushMessage interface {
	IsValid() bool
	ToJson() string
}

type MessageAndroid struct {
	Title         string            `json:"title"`
	Content       string            `json:"content"`
	CustomContent map[string]string `json:"custom_content"`
	AcceptTime    []TimeInterval    `json:"accept_time"`
	BuilderId     int               `json:"builder_id"`
	NId           int               `json:"n_id"`
	Ring          int               `json:"ring"`
	Vibrate       int               `json:"vibrate"`
	Clearable     int               `json:"clearable"`
	Action        ClickAction       `json:"action"`
}

func NewMessageAndroid(title, content string) MessageAndroid {
	var msg MessageAndroid
	msg.Title = title
	msg.Content = content
	msg.CustomContent = make(map[string]string)
	msg.AcceptTime = make([]TimeInterval, 0)
	msg.BuilderId = 0
	msg.NId = 0
	msg.Ring = 0
	msg.Vibrate = 1
	msg.Clearable = 1
	msg.Action.ActionType = TYPE_ACTIVITY
	msg.Action.Brower.ConfirmOnUrl = 1
	return msg
}

func (this MessageAndroid) IsValid() bool {
	if this.Title == "" || this.Content == "" {
		return false
	}
	for _, at := range this.AcceptTime {
		if !at.IsValid() {
			return false
		}
	}
	if this.Ring != 0 && this.Ring != 1 {
		return false
	}
	if this.Vibrate != 0 && this.Vibrate != 1 {
		return false
	}
	if !this.Action.IsValid() {
		return false
	}

	return true
}

func (this MessageAndroid) ToJson() string {
	json, err := json.Marshal(this)
	if err == nil {
		return string(json)
	} else {
		return ""
	}
}

type AlertData struct {
	Body         string   `json:"body"`
	ActionLocKey string   `json:"action-loc-key"`
	LocKey       string   `json:"loc-key"`
	LocArgs      []string `json:"loc-args"`
	LaunchImage  string   `json:"launch-image"`
}

type ApsData struct {
	Alert            AlertData `json:"alert"`
	Badge            int       `json:"badge"`
	Sound            string    `json:"sound"`
	ContentAvailable int       `json:"content-available"`
}

type MessageIOS struct {
	Aps        ApsData        `json:"aps"`
	AcceptTime []TimeInterval `json:"accept_time"`
}

func NewMessageIOS(text string) MessageIOS {
	var msg MessageIOS
	msg.Aps.Alert.Body = text
	return msg
}

func (this MessageIOS) IsValid() bool {
	if this.Aps.Alert.Body == "" {
		return false
	}
	for _, at := range this.AcceptTime {
		if !at.IsValid() {
			return false
		}
	}
	return true
}

func (this MessageIOS) ToJson() string {
	json, err := json.Marshal(this)
	if err == nil {
		return string(json)
	} else {
		return ""
	}
}
