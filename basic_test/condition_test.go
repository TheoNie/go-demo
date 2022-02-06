package basic

import "testing"

func TestIf(t *testing.T) {
	cond := true
	cond1 := true

	//表达式结果必须为bool类型
	if cond {
	} else {
	}

	if cond {
	} else if cond1 {
	} else {
	}

	//可以在判断条件前附加一个赋值语句
	if a := 3; cond {
		t.Log(a)
	}

}

//go中的switch不需要使用break
func TestSwitch(t *testing.T) {

	//将switch当做if的一种特殊写法
	i := 3
	switch {
	case i%2 == 0:
		t.Log("Even")
	case i%2 == 1:
		t.Log("Odd")
	default:
		t.Log("default")
	}

	a := "test1"
	switch a {
	case "test", "test1": //多个结果选项想走同一个分支可以用逗号分隔
		t.Log("test or test1")
	case "test2":
		t.Log("test2")
	default:
		t.Log("unknown")
	}
}
