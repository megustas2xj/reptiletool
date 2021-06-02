package reptiletool

import (
	"encoding/json"
	"os"
)

func GetProxy() (s *Socket,h *AuthSocket){

	file,_:=os.Open("config.json")
	defer file.Close()

	proxyRes:=new(ProxyRes)
	json.NewDecoder(file).Decode(&proxyRes)
	s=&Socket{
		Host: proxyRes.Proxys.Socks.Host,
		Port: proxyRes.Proxys.Socks.Port,
	}
	h=&AuthSocket{
		UserName: proxyRes.Proxys.HTTP.User,
		Password: proxyRes.Proxys.HTTP.Pwd,
		Host: proxyRes.Proxys.HTTP.Host,
		Port: proxyRes.Proxys.HTTP.Port,
	}
	return s,h
}
