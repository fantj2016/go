package dao

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	dbConn *sql.DB
	err error
)
func init()  {
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/user?charset=utf8")
	if err != nil {
		fmt.Printf("init error %v",err)
		panic(err.Error())
	}
}