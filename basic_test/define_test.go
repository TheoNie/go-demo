package basic_test

import (
	"testing"
)

//常量定义
const (
	Monday = iota + 1 //iota是常量计数器，只在常量表达式使用，初始为0,以后每行默认加一，下一行不声明值的常量在上一行基础上自动+1
	Tuesday
	Wednesday
)

func TestConst(t *testing.T) {
	t.Log(Monday)    // print 1
	t.Log(Tuesday)   // print 2
	t.Log(Wednesday) // print 3
}

//变量定义
func TestVarDefine(t *testing.T) {
	var a = 1
	var b int
	b = 2
	c := 3
	var (
		d = 4
		e = 5
	)
	f, g := 6, 7
	t.Log(a, b, c, d, e, f, g)
}
