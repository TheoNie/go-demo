package basic

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"One": 1, "Two": 2, "Three": 3}
	t.Log(m)

	m1 := map[string]int{}
	m1["One"] = 1
	t.Log(m1)

	//make(type, cap)
	m2 := make(map[string]int, 10) //map初始化的时候没有len参数
	t.Log(m2)
}

//访问map里的元素
func TestMapAccess(t *testing.T) {
	m := map[string]int{"Zero": 0, "One": 1, "Two": 2, "Three": 3}

	//map中的元素如果不存在的时候通过下标访问也会默认返回零值，所以需要通过go中访问map时内置的第二个返回值判断其是否真实存在
	if v, ok := m["zero"]; ok {
		t.Log(v)
	} else {
		t.Log("none")
	}
}

func TestMapTravel(t *testing.T) {
	m := map[string]int{"Zero": 0, "One": 1, "Two": 2, "Three": 3}

	for k, v := range m {
		t.Log(k, v)
	}
}

//将函数作为map的值
func TestMapWithFunVal(t *testing.T) {
	m := map[string]func(op int) int{}

	m["one"] = func(op int) int {
		return op
	}
	m["two"] = func(op int) int {
		return 2 * op
	}
	m["three"] = func(op int) int {
		return 3 * op
	}

	t.Log(m["one"](2))   //2
	t.Log(m["two"](2))   //4
	t.Log(m["three"](2)) //6
}
