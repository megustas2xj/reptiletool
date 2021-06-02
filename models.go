package reptiletool

/**
 *
 * @project:
 * @Author: zzw
 * @Date: 2020/07/06 19:25
 * @Description:
 */

type ContentType string
type ProxyType string

const (
	 Json ContentType = "1" // application/json
	 MultipartFromData ContentType = "2" //PostFromData
	 Urlencoded ContentType = "3" //application/x-www-form-urlencoded
	 Socket5 ProxyType = "1"
	 HttpProxy ProxyType = "2"
)

type GetIpRes struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		Host string `json:"Host"`
		UserAgent string `json:"User-Agent"`
		XAmznTraceID string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL string `json:"url"`
}

type AuthSocket struct {
	Host string
	Port string
	UserName string
	Password string
}

type Socket struct {
	Host string
	Port string
}

type ProxyRes struct {
	Proxys struct {
		HTTP struct {
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Pwd string `json:"pwd"`
		} `json:"http"`
		Socks struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"socks"`
	} `json:"proxys"`
}


type ReptileTool struct {

	Debug bool
	Params Params
	ProxyType ProxyType
}

type Params struct {
	Method string
	Url	string
	Header map[string]string
	Body   []byte
	CType ContentType
}

type CallBack struct {
	Content		[]byte
	Err			error
}



