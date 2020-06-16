package reptiletool

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

const ADDRESS = "101.133.159.201"

type ip struct {
	Origin string `json:"origin"`
}

func Getip() string{
	var rsg ip
	ipadd,_,_:=StartRequest("GET","",nil,nil)
	json.Unmarshal(ipadd,&rsg)
	return rsg.Origin
}

func Init() {
	if Getip() != ADDRESS {
		//logger.WriteLog("https","代理设置错误")
		panic("代理设置错误")
		return
	}
}

func PostFromData(bodytype string, httpurl string, headers map[string]string, body map[string]string) (content []byte, statusCode int, err error) {

	Init()

	w := multipart.NewWriter(new(bytes.Buffer))
	for k, v := range body {
		w.WriteField(k, v)
	}
	w.Close()
	postdata, _ := json.Marshal(body)
	headers["Content-Type"] = w.FormDataContentType()
	content, statusCode, err = StartRequest("POST", httpurl, headers, postdata)
	return
}

func StartRequest(method string, httpUrl string, headers map[string]string, body []byte) (content []byte, statusCode int, err error) {

	auth := proxy.Auth{
		User:     "itemb123",
		Password: "kIl8Jl3aKej",
	}

	address := fmt.Sprintf("%s:%s", ADDRESS, "9999")
	dialer, err := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)
	if err != nil {
		return
	}

	//获取ip地址
	if httpUrl == "" {
		httpUrl = "http://httpbin.org/get"
	}
	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequest", err.Error())
		return nil, -1, err
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}
	var resp *http.Response
	httpTransport := &http.Transport{
		//跳过证书验证
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: httpTransport}
	if dialer != nil {
		httpTransport.Dial = dialer.Dial
	}
	resp, err = httpClient.Do(request)
	if err != nil {
		WriteLog("StartRequest", err.Error())
		return nil, -1, err
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	content, err = ioutil.ReadAll(resp.Body)
	return
}

