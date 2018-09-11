package dao

import (
	_"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/jiaofanting/go/go-web/api/defs"
	"github.com/jiaofanting/go/go-web/api/util"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns,err := dbConn.Prepare("insert into users(login_name,password) values(?,?) ")
	if err!=nil {
		log.Println("add遇到错误")
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error)  {
	stmtOut,err := dbConn.Prepare("select password from users where login_name=?")
	if err != nil {
		log.Printf("%s",err)
		return "",err
	}
	var password string
	//执行查询并将结果写入 password 变量中
	err = stmtOut.QueryRow(loginName).Scan(&password)
	//没有 loginName的报错返回就是  ErrNoRows
	if err != nil && err!= sql.ErrNoRows{
		return "",err
	}
	defer stmtOut.Close()
	return password,nil
}

func DeleteUser(loginName string, password string) error {
	stmtDel,err := dbConn.Prepare("delete from users where login_name=? and password=?")
	if err != nil {
		log.Printf("deleteUser error:%s",err)
		return err
	}
	_,err = stmtDel.Exec(loginName,password)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
func AddNewVide(aid int, name string) (*defs.VideoInfo, error) {
	//vid ,err := util.New

}