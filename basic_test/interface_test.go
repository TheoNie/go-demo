package basic

import (
	"sort"
	"testing"
)

//package sort
//type Interface interface {
//	Len() int
//	Less(i, j int) bool // i, j are indices of sequence elements
//	Swap(i, j int)
//}

type CustomType struct {
	a, b, c string
}

func (c CustomType) String() string {
	return c.a + c.b + c.c
}

type MyCustomSort struct {
	s []*CustomType
}

func (x MyCustomSort) Len() int {
	return len(x.s)
}

func (x MyCustomSort) Less(i, j int) bool {
	if x.s[i].a != x.s[j].a {
		return x.s[i].a < x.s[j].a
	}
	if x.s[i].b != x.s[j].b {
		return x.s[i].b < x.s[j].b
	}
	if x.s[i].c != x.s[j].c {
		return x.s[i].c < x.s[j].c
	}
	return false
}

func (x MyCustomSort) Swap(i, j int)  {
	x.s[i], x.s[j] = x.s[j], x.s[i]
}

func TestInterface(t *testing.T) {
	data1 := &CustomType{a: "b", b: "b", c: "c"}
	data2 := &CustomType{a: "a", b: "b", c: "d"}
	data3 := &CustomType{a: "b", b: "b", c: "d"}
	data := []*CustomType{data1, data2, data3}
	t.Log(data)
	sort.Sort(MyCustomSort{data})
	t.Log(data)
}