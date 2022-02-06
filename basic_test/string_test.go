package basic

import (
	"strconv"
	"strings"
	"testing"
)

func TestStrFunction(t *testing.T) {
	//字符串分割
	s := "a,b,c"
	s1 := strings.Split(s, ",")
	t.Log(s1)

	//字符串拼接
	parts := []string{"a", "b", "c"}
	s2 := strings.Join(parts, ",")
	t.Log(s2)
}

func TestStrConv(t *testing.T) {
	//整型转换为字符串
	s := strconv.Itoa(10)
	t.Log("str:" + s)

	//字符串转换为整型
	if v, err := strconv.Atoi("xx"); nil == err {
		t.Log("number is ", v)
	} else {
		t.Log("error", err)
	}
}

//rune用来表示Unicode的code point,本质是int32
func TestStrToRune(t *testing.T) {
	s := "语言"
	//打印出每个字符和其对应的unicode,其中的range会自动把遍历的字符串转为rune
	for _, v := range s {
		t.Logf("%[1]c %[1]x", v)
	}

	c := []rune(s)
	t.Log(c)
	t.Logf("语的unicode, %x", c[0])
	t.Logf("语的UTF-8, %x", "语")
}
