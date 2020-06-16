package reptiletool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var connect *sql.DB

const (
	DB_CHARSET = "utf8"
)

func InitDb(Username,Password,DBName string) {
	dns:=fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=%s",Username,Password,DBName,DB_CHARSET)
	connect, _ = sql.Open("mysql", dns)
	connect.SetMaxOpenConns(128)
	connect.SetMaxIdleConns(32)
	connect.SetConnMaxLifetime(120 * time.Second)
	err := connect.Ping()
	if err != nil {
		WriteLog("error", err.Error())
		panic(err)
	}
}

func Insert(Username,Password,DBName,tableName,uid string) {
	InitDb(Username,Password,DBName)
	tx, err1 := connect.Begin()
	if err1 != nil {
		panic(err1)
	}
	stmt, _ := connect.Prepare(fmt.Sprintf("INSERT ignore into %s (uid) values (?)", tableName))
	_, err := stmt.Exec(uid)
	WriteLog(fmt.Sprintf("%s", tableName), fmt.Sprintf("%s", uid))
	if err != nil {
		WriteLog("error", err.Error())
		tx.Rollback()
	}
	tx.Commit()
	connect.Close()
}

func InsertTx(Username,Password,DBName,tableName,uid string) {
	InitDb(Username,Password,DBName)
	tx, err1 := connect.Begin()
	if err1 != nil {
		panic(err1)
	}
	tx.Exec(fmt.Sprintf("INSERT ignore into %s (uid) values (?)", tableName), uid)
	tx.Commit()
}
