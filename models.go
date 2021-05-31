package reptiletool

/**
 *
 * @project:
 * @Author: zzw
 * @Date: 2020/07/06 19:25
 * @Description:
 */


type ConfigJson struct {
	Database struct {
		Mysql struct {
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Pwd string `json:"pwd"`
			Name string `json:"name"`
			Driver string `json:"driver"`
		} `json:"mysql"`
		DBCHARSET string `json:"DB_CHARSET"`
	} `json:"database"`
	Proxys struct {
		Proxy struct {
			Host string `json:"host"`
			Port string `json:"port"`
			User string `json:"user"`
			Pwd string `json:"pwd"`
		} `json:"proxy"`
		Socks struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"socks"`
	} `json:"proxys"`
}

type ConfigDb struct {
	Host 	  string
	Port 	  string
	User 	  string
	Pwd  	  string
	Name 	  string
	Driver    string
	DbCharset string
}

type Proxy struct {

	Host 	  string
	Port 	  string
	User 	  string
	Pwd  	  string
	SocksHost string
	SocksPort string

}


type Params struct {

	Method string
	Header map[string]string


}

type CallBack struct {
	content		[]byte
	err			error
}



