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
	"os"
)

var callBack *CallBack

func GetProxy() *Proxy{

	file,_:=os.Open("config.json")
	defer file.Close()

	configJson:=ConfigJson{}
	json.NewDecoder(file).Decode(&configJson)

	config:=Proxy{
		Host:configJson.Proxys.Proxy.Host,
		Port:configJson.Proxys.Proxy.Port,
		User:configJson.Proxys.Proxy.User,
		Pwd:configJson.Proxys.Proxy.Pwd,
		SocksHost:configJson.Proxys.Socks.Host,
		SocksPort:configJson.Proxys.Socks.Port,
	}
	return &config
}

func init(){
	callBack=new(CallBack)
}

func PostFromData(bodytype string, httpurl string, headers map[string]string, body map[string]string) *CallBack {

	w := multipart.NewWriter(new(bytes.Buffer))
	for k, v := range body {
		w.WriteField(k, v)
	}
	w.Close()
	postdata, _ := json.Marshal(body)
	headers["Content-Type"] = w.FormDataContentType()
	callBack= StartRequest("POST", httpurl, headers, postdata)
	return callBack
}

func StartRequest(method string, httpUrl string, headers map[string]string, body []byte) *CallBack {

	value:=GetProxy()

	auth := proxy.Auth{
		User:     value.User,
		Password: value.Pwd,
	}

	address := fmt.Sprintf("%s:%s", value.Host, value.Port)
	dialer, err := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)
	if err != nil {
		callBack.err=err
		return callBack
	}

	//获取ip地址
	if httpUrl == "" {
		httpUrl = "http://httpbin.org/get"
	}
	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequest", err.Error())
		callBack.err=err
		return callBack
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
		callBack.err=err
		return callBack
	}
	defer resp.Body.Close()
	callBack.content, callBack.err= ioutil.ReadAll(resp.Body)
	return callBack
}


func StartRequestNoProxy(method string, httpUrl string, headers map[string]string, body []byte) *CallBack {

	//获取ip地址
	if httpUrl == "" {
		httpUrl = "http://httpbin.org/get"
	}
	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequestNoProxy", err.Error())
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	return
}

func StartRequestSocks(method string, httpUrl string, headers map[string]string, body []byte) *CallBack {

	value:=GetProxy()

	dialer, err := proxy.SOCKS5("tcp",fmt.Sprintf("%v:%v",value.SocksHost,value.SocksPort),nil, proxy.Direct)
	if err != nil {
		return
	}

	request, err := http.NewRequest(method, httpUrl, bytes.NewBuffer(body))
	if err != nil {
		WriteLog("StartRequestSocks", err.Error())
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	return
}
