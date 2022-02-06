package basic

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(innerFunc func(op int) int) func(int) int {
	return func(i int) int {
		start := time.Now()
		result := innerFunc(i)
		end := time.Now()
		fmt.Println("timeSpent is", end.Sub(start))
		return result
	}
}

func func1(op int) int {
	time.Sleep(100)
	return op
}

//利用函数可以作为方法入参的特性，计算函数执行花费时间
func TestFunc(t *testing.T) {
	func14TimeSpent := timeSpent(func1)
	result := func14TimeSpent(2)
	t.Log(result)
}
