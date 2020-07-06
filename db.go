package reptiletool

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)


var connect *sql.DB

func GetConfig() *ConfigDb{

	file,_:=os.Open("config.json")
	defer file.Close()

	configJson:=ConfigJson{}
	json.NewDecoder(file).Decode(&configJson)

	config:=ConfigDb{
		Host:configJson.Database.Mysql.Host,
		Port:configJson.Database.Mysql.Port,
		User:configJson.Database.Mysql.User,
		Pwd:configJson.Database.Mysql.Pwd,
		Name:configJson.Database.Mysql.Name,
		Driver:configJson.Database.Mysql.Driver,
		DbCharset:configJson.Database.DBCHARSET,
	}
	return &config
}

func InitDb() {

	value:=GetConfig()
	dns:=fmt.Sprintf("%s:%s@tcp(%v:%v)/%s?charset=%s",value.User,value.Pwd,value.Host,value.Port,value.Name,value.DbCharset)
	connect, _ = sql.Open("mysql", dns)
}

func NewConnection() *sql.DB{

	InitDb()
	return connect
}

func InsertUid(tableName,uid string) error{

	InitDb()
	stmt,_:=connect.Prepare(fmt.Sprintf("INSERT ignore into %s (uid) values (?)", tableName))
	_,err:=stmt.Exec(uid)
	return err
}
