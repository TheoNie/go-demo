package basic

import "testing"

func TestSliceInit(t *testing.T) {
	var s []int //与数组不同的是，切片初始化时中括号中不需要添加数字或省略号
	s = append(s, 1, 2)
	t.Log(s)

	s1 := []int{1, 2, 3}
	t.Log(s1)

	//make([]type, len, cap), len个元素会有默认值，cap代表容量，不可以越界访问未初始化元素
	s2 := make([]int, 1, 3)
	t.Log(len(s2), cap(s2)) //1,3
	s2 = append(s2, 4, 5, 6, 7)
	t.Log(s2)
	t.Log(len(s2), cap(s2)) //5, 6  在append后，元素个数增加，而cap因为不够容纳也会自动扩容,每次扩容为前一次cap的2倍
}

func TestSliceShareMemory(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	s1 := s[1:3]
	s2 := s[2:3]

	//截取之后的切片，len会等于截取长度的大小，而cap会等于，被截取的切片从截取起始位置到该切片末尾的长度大小
	t.Log(len(s1), cap(s1)) //2,5
	t.Log(len(s2), cap(s2)) //1,4

	//改变任意元素的大小，都会影响到相关的原始切片和同源切片
	s2[0] = 0
	t.Log(s)  //1 2 0 4 5 6
	t.Log(s1) //2 0
	t.Log(s2) //0
}
