package dao

import (
	_"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/jiaofanting/go/go-web/api/defs"
	"github.com/satori/go.uuid"
	"time"
	"fmt"
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
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	uuid,err := uuid.NewV4()
	var vid  = uuid.String()
	if err != nil {
		return nil,err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns,err := dbConn.Prepare(`insert into video_info
			(id,author_id,name,display_ctime) values(?,?,?,?)`)
	if err != nil {
		return nil,err
	}
	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err != nil {
		fmt.Println("这里有错")
		return nil,err
	}
	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
func AddNewComments(vid string, aid int, content string) error {
	uuid, err := uuid.NewV4()
	var id = uuid.String()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	var res [] *defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	/**
	 将结果放入数组
	 */
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}