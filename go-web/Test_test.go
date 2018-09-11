package main

import (
	"testing"
	"fmt"
)

/**
1. 必须import testing
2. 文件名必须是 xxx_test.go
3. 测试性能传值 *testing.B
4. t.SkipNow()跳过当前test，直接按pass处理
 */
func TestPrint(t *testing.T)  {
	//res := Print1to20()

	//if res != 210{
	//	t.Error("return value not valid")
	//}

	t.Run("a1", func(t *testing.T) {
		fmt.Print("a1")
	})
	t.Run("test1",test1)
}
func test1(t *testing.T) {
	fmt.Println("test1")
}

/**
test之 benchmark
一般会跑b.N次,注意
cmd: go test -bench=.
3000            425596 ns/op
PASS
ok      _/D_/workspace/learn-src/go-learn/go-web        1.527s
 */

func BenchmarkAll(b *testing.B)  {
	for n:= 0 ;n<b.N;n++{
		fmt.Println("benchmark")
	}

}