package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	condition1()
	fmt.Println(
		converToBin(2),
		converToBin(4),
	)
	pringFile("hello.txt")
}
func condition1()  {
	/**
	if 用法1：
	 */
	const filename = "hello.txt"
	contents,err := ioutil.ReadFile(filename)
	if err!= nil {
		fmt.Print(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
	/**
	if 用法2：
	 */
	if contents,err := ioutil.ReadFile(filename);err != nil{
		fmt.Print(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
}
/**
switch用法:默认自动break
 */
func condition2(a,b int,op string) int {
	var result int
	switch op {
	case "+":
		result = a+b
	case "-":
		result = a-b
	default:
		panic(op)
	}
	return  result
}
/**
for 循环:十进制转换二进制
 */
func converToBin(n int) string {
	result := ""
	for ;n>0 ;n /= 2  {
		lsb := n%2
		result = strconv.Itoa(lsb) + result //strconv.Itoa(lsb)为强制转换成string
	}
	return result
}
/**
for语法   for{} 相当于while，go没有while
 */
func pringFile(filename string)  {
	file,err  := os.Open(filename)
	if err != nil {
		//panic 终止程序运行
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}



