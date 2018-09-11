package dao

import (
	"testing"
	"fmt"
)
func TestMain(m *testing.M) {
	m.Run()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("add",testAddUserCredential)
	t.Run("get",testGetUserCredential)
	t.Run("delete",testDeleteUser)
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
