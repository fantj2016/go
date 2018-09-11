package dao

import (
	"testing"
	"fmt"
	"strconv"
	"time"
)
var tempvid string
func TestMain(m *testing.M) {
	m.Run()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("add",testAddUserCredential)
	t.Run("get",testGetUserCredential)
	t.Run("delete",testDeleteUser)
}
func TestVideoWorkFlow(t *testing.T) {
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("fantj","root")
	if err != nil {
		fmt.Println("add test 错误")
		t.Errorf("Error of AddUser:%v",err)
	}
}

func testGetUserCredential(t *testing.T) {
	password,err := GetUserCredential("fantj")
	if err != nil {
		t.Errorf("getuser error:%v",err)
	}
	fmt.Print(password)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("fantj","root")
	if err != nil {
		t.Errorf("deleteuser error:%v",err)
	}

}
func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}
func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}
func TestComments(t *testing.T) {
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "1"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "1"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
