package main

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE="c"
)

type controlChan chan string
//interface 表示泛型
type dataChan chan interface{}
//用fn定义了一个函数，参数 dc，返回error
type fn func(dc dataChan) error

