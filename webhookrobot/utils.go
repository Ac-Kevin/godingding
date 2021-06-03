package webhookrobot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

/*HTTPBaseRequest 发送http请求
传入参数：url 请求主机地址
				 params 参数
				 contentType 参数格式
				 method 请求方法 "POST" OR "GET" 注意要大写
				 timeout 超时设置 单位 s
				 headers map[string]string 非必填参数 | 可用于添加实际用户IP地址 例子：map[string]string{"Remote_addr": "用户IP地址"}
*/
func HTTPBaseRequest(url, params, contentType, method string, timeout int, headers ...map[string]string) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	var req *http.Request
	var err error
	switch method {
	case "POST":
		req, err = http.NewRequest("POST", url, strings.NewReader(params))
		if contentType == "" {
			contentType = "application/x-www-form-urlencoded"
		}
		req.Header.Set("Content-Type", contentType)
	case "GET":
		req, err = http.NewRequest("GET", url+"?"+params, nil)
	default:
		err = fmt.Errorf("不可识别method：'%s'", method)
	}
	if err != nil {
		return "", fmt.Errorf("HTTP-Request-Err :%s", err)
	}
	//设置heads
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Add(key, value)
		}
	}
	//http请求
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP-Request-Err :%s", err)
	}
	//读取返回信息
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	//添加404错误检查
	bodyString := string(body)
	if len(bodyString) > 17 && bodyString[:18] == "404 page not found" {
		return bodyString, errors.New("404 page not found")
	}
	return bodyString, nil
}
