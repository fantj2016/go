package main

import "fmt"

/**
常量定义和枚举
 */
func main() {
	consts()
	enums()
}
/**
 * 常量定义
 */
func consts(){
	const(
		name  = "fantj"
		age = 21
	)
	const firstName  = "FantJ"
	fmt.Println(name,age,firstName)
}
/**
enum 用法1
 */
func enums()  {
	const(
		name = "fantj"
		age = 21
		javaAge = 2
	)
	fmt.Println(name,age,javaAge)
	/**
	enum 用法2
	 */
	const(
		num_1 = iota  //iota表示 从当前数开始从0开始自增
		num_2
		num_3
	)
	fmt.Print(num_1,num_2,num_3)
}
