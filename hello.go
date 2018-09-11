package main

import (
	"fmt"
)

func variable()  {
	var a int
	var b string
	fmt.Printf("%d %q\n",a,b)
}

func main() {
	fmt.Print("hello world")
	variable()
	consts()
	enums()
}
