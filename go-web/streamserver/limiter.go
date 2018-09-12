package main

import (
	"log"
)

/**
流量控制： bucket token 算法

一个request 进来，发一个token，带着token才能请求
go 语言用共享通道来实现共享，而不是共享内存（因为共享内存必须加锁，影响性能）
 */

 type ConnLimiter struct {
 	concurrentConn int
 	bucket chan int
 }
//构造器必须自己写
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),
	}
}
//获取连接
func (cl *ConnLimiter) getConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation")
		return false
	}
	//赋值1
	cl.bucket <- 1
	return true
}
//施放连接
func (cl *ConnLimiter) ReleaseConn()  {
	c := <-cl.bucket
	log.Printf("New connction coming: %d",c)
}