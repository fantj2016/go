package main

type treeNode struct {
	value int
	left,right *treeNode
}
func main() {
	/**
	go语言仅支持封装，不支持继承和多台
	go语言灭有class，只有struct
	 */
	 root := treeNode{value:3}
	 root.left = &treeNode{}
	 root.right = &treeNode{5,nil,nil}
	 root.right.left = new(treeNode)
}
