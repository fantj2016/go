package main

import "fmt"

func main() {
	variable1()
}


func variable1()  {
	/**
	第一种，指定变量类型，声明后若不赋值，使用默认值。
	 */
	var v_name1 string
	v_name1 = "fantj"
	fmt.Println(v_name1)
	/**
	根据值自行判定变量类型。
	 */
	var v_name2 = "fantj"
	fmt.Println(v_name2)
	/**
	 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
	 */
	name := "fantj"
	fmt.Println(name)
	/**
	多变量声明
	 */
	 var a,b,c = 1,2,3
	 fmt.Println(a,b,c)
}
