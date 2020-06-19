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
	"github.com/unknwon/goconfig"
)

var ADDRESS string

type ip struct {
	Origin string `json:"origin"`
}

func Init(){

	config, _ := goconfig.LoadConfigFile("conf.ini")
	ADDRESS,_=config.GetValue("header","nonce")

}

func PostFromData(bodytype string, httpurl string, headers map[string]string, body map[string]string) (content []byte, statusCode int, err error) {

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


func StartRequestNoProxy(method string, httpUrl string, headers map[string]string, body []byte) (content []byte, statusCode int, err error) {

	//获取ip地址
	if httpUrl == "" {
		httpUrl = "http://httpbin.org/get"
	}
	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequestNoProxy", err.Error())
		return nil, -1, err
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}
	var resp *http.Response

	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		WriteLog("StartRequestNoProxy", err.Error())
		return nil, -1, err
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	content, err = ioutil.ReadAll(resp.Body)
	return
}

func StartRequestSocks(method string, httpUrl string, headers map[string]string, body []byte,port int) (content []byte, statusCode int, err error) {

	dialer, err := proxy.SOCKS5("tcp",fmt.Sprintf("127.0.0.1:%v",port),nil, proxy.Direct)
	if err != nil {
		return
	}

	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequestSocks", err.Error())
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
		WriteLog("StartRequestSocks", err.Error())
		return nil, -1, err
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	content, err = ioutil.ReadAll(resp.Body)
	return
}
