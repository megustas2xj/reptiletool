package reptiletool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var connect *sql.DB

func InitDb() {
	connect, _ = sql.Open("mysql", "root:zzw198419@tcp(127.0.0.1:3306)/admin?charset=utf8")
	connect.SetMaxOpenConns(128)
	connect.SetMaxIdleConns(32)
	connect.SetConnMaxLifetime(120 * time.Second)
	err := connect.Ping()
	if err != nil {
		WriteLog("error", err.Error())
		panic(err)
	}
}

func Insert(tableName string, uid string) {
	InitDb()
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

func InsertTx(tablename string, uid string) {
	InitDb()
	tx, err1 := connect.Begin()
	if err1 != nil {
		panic(err1)
	}
	tx.Exec(fmt.Sprintf("INSERT ignore into %s (uid) values (?)", tablename), uid)
	tx.Commit()
}
