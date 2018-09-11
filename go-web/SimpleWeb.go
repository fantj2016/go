package main

import (
	"net/http"
	"io"
)

func Print1to20() int {
	res := 0
	for i:=1 ;i<=20;i++{
		res += 1
	}
	return res
}
func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"hello go")
}

func main() {
	http.HandleFunc("/",firstPage)
	http.ListenAndServe(":8080",nil)
}

