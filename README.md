## reptiletool介绍
reptiletool 是一个http请求的工具

### 下载
go get github.com/megustas2xj/reptiletool

### 说明
自行修改config.json 配置文件

# example

    import (
        "github.com/megustas2xj/reptiletool"
        "log"
    )
## 1.Using Proxy
    func main() {

        httpUrl:= "http://httpbin.org/get"
        r:=new(reptiletool.ReptileTool)

        p:=reptiletool.ProxyConfig{
            Host: "127.0.0.1",
            Port: "10808",
        }

        c:=r.StartRequest("GET",httpUrl,nil,nil).SetProxy(reptiletool.Socket5,p).SetDebug(true).MakeRequest()
        if c.Content!=nil{
            log.Printf("%s",string(c.Content))
        }

    }

## 2.Not using a proxy

	httpUrl:= "http://httpbin.org/get"
	r:=new(reptiletool.ReptileTool)

	c:=r.StartRequest("GET",httpUrl,nil,nil).SetDebug(true).MakeRequest()
	if c.Content!=nil{
		log.Printf("%s",string(c.Content))
	}