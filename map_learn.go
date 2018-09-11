package main

import "fmt"

/**
map[key类型]value类型{}
 */
func main() {
	m := map[string]string{
		"name":"fantj",
		"school":"tyut",
	}
	m2 := make(map[string]int)
	fmt.Print(m,m2)
	/**
	for遍历 key-value
	 */
	for k,v := range m{
		fmt.Println(k,v)
	}
	/**
	从key中获取value
	 */
	 name := m["name"]
	 fmt.Println(name)
	
}
