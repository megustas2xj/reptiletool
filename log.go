package reptiletool

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	//LOGPATH  LOGPATH/time.Now().Format(FORMAT)/*.log
	LOGPATH = "log/"
	//FORMAT .
	FORMAT = "20060102"
	//LineFeed 换行
	LineFeed = "\r\n"
)

//以天为基准,存日志
var path = LOGPATH + time.Now().Format(FORMAT) + "/"

//WriteLog return error
func WriteLog(fileName string, msg string) error {
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	msg="["+fmt.Sprintf("%v",time.Now().Format("2006-01-02 15:04:05.000"))+"]: "+msg
	_, err = io.WriteString(f,msg+LineFeed)

	defer f.Close()
	return err
}

func WriteId(msg string) error{
	if !IsExist(path) {
		return CreateDir(path)
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+"idList.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	_, err = io.WriteString(f,msg+LineFeed)
	defer f.Close()
	return err
}

//CreateDir  文件夹创建
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}