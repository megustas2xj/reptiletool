package reptiletool

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

var c *CallBack
func init(){
	c=new(CallBack)
}

func (r *ReptileTool)StartRequest(method string,url string,header map[string]string,body []byte)*ReptileTool{

	r.Params.Method=method
	r.Params.Header=header
	r.Params.Body = body
	r.Params.Url = url
	return r
}

func (r *ReptileTool)SetContentType(contentType ContentType) *ReptileTool{
	r.Params.CType=contentType
	return r
}

func (r *ReptileTool)SetProxy(proxyType ProxyType,config ProxyConfig) *ReptileTool{
	r.ProxyType=proxyType
	r.ProxyConfig=config
	return r
}

func (r *ReptileTool)SetDebug(flag bool) *ReptileTool{
	r.Debug=flag
	return r
}

func (r *ReptileTool)MakeRequest() *CallBack{

	var body map[string]string
	if r.Params.Body!=nil{
		body=GetKeyValue(r.Params.Body)
	}
	var postData []byte
	if r.Params.CType==""{
		r.Params.CType=Json
	}
	cp :=r.Params.CType
	switch cp {
	case MultipartFromData:
		w := multipart.NewWriter(new(bytes.Buffer))
		for k, v := range body {
			w.WriteField(k, v)
		}
		w.Close()
		postData, _ = json.Marshal(body)
		r.Params.Body=postData
		r.Params.Header["Content-Type"] = w.FormDataContentType()
	case Urlencoded:
		DataUlrval:=url.Values{}
		for s, s2 := range body {
			DataUlrval.Add(s,s2)
		}
		r.Params.Body=[]byte(DataUlrval.Encode())
	}

	if r.ProxyType!=""{
		return StartRequestProxy(r)
	}
	return StartRequest(r)
}

func StartRequestProxy(r *ReptileTool) *CallBack{

	var auth proxy.Auth
	var address string
	proxyType := r.ProxyType
	switch proxyType {
	case Socket5:
		address = fmt.Sprintf("%s:%s", r.ProxyConfig.Host, r.ProxyConfig.Port)
	case HttpProxy:
		auth=proxy.Auth{
			User:     r.ProxyConfig.User,
			Password: r.ProxyConfig.Pwd,
		}
		address = fmt.Sprintf("%s:%s", r.ProxyConfig.Host,r.ProxyConfig.Port)
	}

	dialer, err := proxy.SOCKS5("tcp", address, &auth, proxy.Direct)
	if err != nil {
		if r.Debug{
			c.Err=err
			log.Printf("[error] %s %s",r.Params.Url,err.Error())
		}
		return c
	}
	request, err := http.NewRequest(r.Params.Method, r.Params.Url, bytes.NewBuffer(r.Params.Body))
	if err != nil {
		if r.Debug{
			c.Err=err
			log.Printf("[error] %s %s",r.Params.Url,err.Error())
		}
		return c
	}
	if r.Params.Header != nil {
		for key, val := range r.Params.Header {
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
		if r.Debug{
			c.Err=err
			log.Printf("[error] %s %s",r.Params.Url,err.Error())
		}
		return c
	}

	defer resp.Body.Close()
	content, err:= ioutil.ReadAll(resp.Body)
	c.Content=content
	return c

}

func StartRequest(r *ReptileTool) *CallBack {

	request, err := http.NewRequest(r.Params.Method, r.Params.Url, bytes.NewBuffer(r.Params.Body))
	if err != nil {
		if r.Debug {
			log.Printf("[error] %s  %s",r.Params.Url,err.Error())
			c.Err=err
			return c
		}
	}
	if r.Params.Header != nil {
		for key, val := range r.Params.Header {
			request.Header.Add(key, val)
		}
	}
	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		if r.Debug{
			log.Printf("[error] %s %s",r.Params.Url,err.Error())
			c.Err=err
			return c
		}
	}
	defer resp.Body.Close()
	content, _:= ioutil.ReadAll(resp.Body)
	c.Content=content
	return c
}

func CheckIP(){

	httpUrl:= "http://httpbin.org/get"
	r:=new(ReptileTool)
	config:=ProxyConfig{Host: "127.0.0.1",Port: "10808"}
	c:=r.StartRequest("GET",httpUrl,nil,nil).SetProxy(Socket5,config).SetDebug(true).MakeRequest()
	ipRes:=new(GetIpRes)
	if c.Content!=nil{
		json.Unmarshal(c.Content,&ipRes)
		log.Printf("【originIp】 %s",ipRes.Origin)
	}
}