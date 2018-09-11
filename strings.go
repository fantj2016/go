package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网" //每个中文3字节
	fmt.Print(len(s))
	bytes := []byte(s)
	for len(bytes)>0  {
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("%c",ch)
	}
	fmt.Println()
}
