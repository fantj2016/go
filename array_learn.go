package main

import "fmt"

func main() {
	var array1 [5] int
	array2 := [3]int{1,2,3}
	array3 := [...]int{1,2,3}
	var grid [4][5]int
	fmt.Print(
		array1,array2,array3,grid,
	)
	/**
	遍历数组
	 */
	 for i := range array3{
	 	fmt.Println(array3[i])
	 }
	 /**
	  索引值 + value 形式遍历
	  */
	for i,v := range array3{
		fmt.Println(i,v)
	}
	/**
	只遍历值
	 */
	 for _,v := range  array3{
	 	array3[0] = 100
	 	fmt.Println(v)
	 }
}
func array_learn()  {

}

func iteriable()  {

}
