package request

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/maczh/mgin/logs"
	"github.com/maczh/mgin/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url, token string, data interface{}) (string, error) {
	logs.Debug("请求参数：{}", data)
	dataStr := utils.ToJSON(data)
	obj := make(map[string]interface{})
	utils.FromJSON(dataStr, &obj)

	paramStr := ""
	for k, v := range obj {
		if IsList(v) {
			list := v.([]interface{})
			for _, s := range list {
				paramStr += fmt.Sprintf("%s=%s&", k, fmt.Sprint(s))
			}
		} else {
			switch v.(type) {
			case float64:
				paramStr += fmt.Sprintf("%s=%f&", k, v)
			case string:
				paramStr += fmt.Sprintf("%s=%s&", k, v)
			default:
				paramStr += fmt.Sprintf("%s=%v&", k, v)
			}
		}
	}
	if len(paramStr) > 0 {
		paramStr = string([]byte(paramStr)[:len(paramStr)-1]) //去除多余的&
		url = url + "?" + paramStr
	}
	logs.Debug("url:{}", url)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr} //忽略https的请求
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Auth-Token", token)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("解析错误:{}", err.Error())
		return "", err
	}
	result := string(body)
	logs.Error("http请求返回的数据:{}", result)
	return result, nil
}

func Post(url, token string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	logs.Debug("请求url：{}", url)
	dataStr := utils.ToJSON(params)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr} //忽略https的请求

	req, err := http.NewRequest("POST", url, strings.NewReader(dataStr))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Auth-Token", token)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("解析错误:{}", err.Error())
		return "", err
	}
	resStr := string(body)
	logs.Debug("http请求返回的数据:{}", resStr)
	if res.StatusCode == 202 {
		return resStr, nil
	}
	if res.StatusCode != 200 {
		return "", errors.New(resStr)
	}
	return resStr, nil
}

func Delete(url, token string, params interface{}) (string, error) {
	logs.Debug("请求参数：{}", params)
	logs.Debug("请求url：{}", url)
	dataStr := utils.ToJSON(params)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr} //忽略https的请求

	req, err := http.NewRequest("DELETE", url, strings.NewReader(dataStr))
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Auth-Token", token)
	res, err := client.Do(req)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error("解析错误:{}", err.Error())
		return "", err
	}
	resStr := string(body)
	logs.Debug("http请求返回的数据:{}", resStr)
	if res.StatusCode == 202 {
		return "", nil
	}
	if res.StatusCode != 200 {
		return "", errors.New(resStr)
	}
	return resStr, nil
}

func GetToken(username, password, projectName string) string {
	url := "https://iam-apigateway-proxy.hyy.com:5443/v3/auth/tokens"
	method := "POST"

	payload := fmt.Sprintf(`{"auth": {"identity": {"methods": ["password"],"password": {"user": {"domain": {"name": "%s"},"name": "%s","password": "%s"}}},"scope": {"project": {"name": "%s"}}}}`, username, username, password, projectName)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr} //忽略https的请求
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		logs.Error("http请求错误:{}", err.Error())
		return ""
	}
	defer res.Body.Close()
	respHeader := res.Header
	token := respHeader.Get("X-Subject-Token")

	return token
}

func IsList(params interface{}) bool {
	switch params.(type) {
	case []interface{}:
		return true
	}
	return false
}
