package goqqpushapi

type Result struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_message"`
	Data    interface{}
}
