package webhookrobot

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// WebHookRot def .
type WebHookRot struct {
	sendurl   string
	secretKey string
	timeout   int
}

type WebHookRotOption struct {
	SendUrl   string
	SecretKey string
	TimeOut   int
}

// NewWebHookRot create a rot
func NewWebHookRot(option WebHookRotOption) *WebHookRot {
	var rot = &WebHookRot{sendurl: option.SendUrl, secretKey: option.SecretKey, timeout: 10}
	if option.TimeOut > 5 {
		rot.timeout = option.TimeOut
	}
	return rot
}

// SendMsg send a msg
func (w *WebHookRot) SendMsg(m Msg) (res Response, err error) {
	var urlpath = w.sendurl
	if w.secretKey != "" {
		var timestamp = strconv.FormatInt(time.Now().Unix()*1000, 10)
		var sign = w.sign(timestamp)
		urlpath = urlpath + fmt.Sprintf("&timestamp=%s&sign=%s", timestamp, url.QueryEscape(sign))
	}
	return w.request(urlpath, m.ToJSONString())
}

func (w *WebHookRot) sign(timestamp string) string {
	return HmacSha256(fmt.Sprintf("%s\n%s", timestamp, w.secretKey), w.secretKey)
}
