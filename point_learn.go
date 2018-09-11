package main

import "fmt"

func main() {
	point()
	a := 3
	//值传递结果
	passByVal(a)
	fmt.Print(a)
	//指针传递结果
	passByRef(&a)
	fmt.Print(a)
}

func point()  {
	var a int = 2
	var pa *int =&a
	*pa = 3
	fmt.Print(a)
}
func passByVal( a int)  {
	a++
}
func passByRef(a *int)  {
	*a++
}
