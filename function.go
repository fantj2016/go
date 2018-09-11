package main

import "fmt"

/*
函数学习
func 函数名(参数...)(返回值...){

}
 */
func main() {
	fmt.Print(count(10,5))
	pring(count2,10,5)
	fmt.Println("string",pring(count2,10,5))
}
/**
返回加减法
 */
func count(a, b int) (int,int) {
	return a+b,a-b
}
func count2(a, b int) (int) {
	return a+b
}
/**
函数做参数
 */
func pring(count func(int,int) int,a,b int) int {
	fmt.Println(count(10,5))
	return count(a,b)
}