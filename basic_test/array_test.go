package basic

import "testing"

//数组声明
func TestArrayDefine(t *testing.T) {
	var a [3]int //声明一个数组，默认元素都为0
	a[1] = 2

	b := [3]int{1, 2, 3} //声明并直接初始化

	c := [2][2]int{{1, 2}, {1, 2}} //声明并直接初始化多维数组

	d := [...]int{1, 2, 3} //省略具体数组元素个数的初始化

	t.Log(a, b, c, d)
}

//数组遍历
func TestArrayTraverse(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}

	//普通的遍历
	for i := 0; i < len(a); i++ {
		t.Log(a[i])
	}

	//foreach类型的遍历,_在go中为缺省值，此处代表range返回的第一个值，也就是数组索引，_可以替换为变量，供下面语句使用
	for _, v := range a {
		t.Log(v)
	}
}

//数组的截取
func TestArrayIntercept(t *testing.T) {
	//array[开始索引(包含),结束索引(不包含)]
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[1:2]) //打印2
	t.Log(a[1:3]) //打印2，3
	t.Log(a[1:])  //打印2,3,4,5
	t.Log(a[1:])  //打印2,3,4,5
	t.Log(a[:])   //打印1,2,3,4,5
}
