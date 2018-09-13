package main

import "fmt"
/**
slice 是对数据底层的一个view
slice 可以向后扩展，但是不能像前扩展
 */
func main() {
	/**
	利用 slice 获取数组中指定索引之间的值
	 */
/*	var arr = [...]int{0,1,2,3,4,5,6,7,8}
	s := arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[2:])  //获取 索引值在2以后的数
	fmt.Println(arr[:6])  //获取 所引致在6以前的数
	fmt.Println(arr[:])   //获取 所有数*/

	createSlice()
}

func createSlice()  {
	s := make([]int,16)
	s2 := make([]int,16,32)
	fmt.Println(s,s2)
	/**
	复制slice
	 */
	copy(s2,s)
	/**
	删掉slice:假设删掉s的第三个元素
	 */
	 s = append(s[:3],s[4:]...)
	 /**
	 取头值/取尾值
	  */
	  front := s[0];
	  tail := s[len(s)-1];
	  fmt.Println(front,tail);
}

