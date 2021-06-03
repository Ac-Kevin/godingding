package webhookrobot

import (
	"errors"
	"testing"
)

func TestSendText(t *testing.T) {
	var sendurl = `https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxxxx`
	var SecretKey = `SECxxxxxxxxxxxxxxxxxx`
	rot := NewWebHookRot(WebHookRotOption{SendUrl: sendurl, SecretKey: SecretKey})
	var msg = Msg{}
	msg.SetText(MsgText{Content: "AV8D 嗨嗨嗨~"}, MsgAt{})
	res, err := rot.SendMsg(msg)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	if res.Errcode != 0 {
		t.Error(errors.New(res.Errmsg))
		t.Fail()
	}
}
