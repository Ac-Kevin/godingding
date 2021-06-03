package webhookrobot

import (
	"encoding/json"
)

// Response def .
type Response struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (w *WebHookRot) request(url, parameter string) (res Response, err error) {
	result, err := HTTPBaseRequest(url, parameter, "application/json", "POST", w.timeout)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(result), &res)
	return
}
