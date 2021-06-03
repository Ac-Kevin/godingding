# Go DD
```
package main

import (
	"github.com/Ac-Kevin/godingding/webhookrobot"
)

func main() {
	var SendUrl = `https://oapi.dingtalk.com/robot/send?access_token=xxxxxx`
	var SecretKey = `you'r secret key`
	robot := webhookrobot.NewWebHookRot(webhookrobot.WebHookRotOption{SendUrl: SendUrl, SecretKey: SecretKey})
	var msg webhookrobot.Msg
	msg.SetText(webhookrobot.MsgText{Content: "Hi~"}, webhookrobot.MsgAt{})
	robot.SendMsg(msg)
}

```
