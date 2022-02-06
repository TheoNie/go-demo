package basic

import (
	"fmt"
	"testing"
)

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	FreezingC Celsius = 0   // 结冰点温度
	BoilingC  Celsius = 100 // 沸水温度
)

func TestType(t *testing.T) {
	t.Log(BoilingC - FreezingC)
	boilingF := cToF(BoilingC)
	t.Log(boilingF - cToF(FreezingC))
	//t.Log(boilingF - FreezingC)  会出现类型不匹配的错误
	t.Log(BoilingC)
}

func cToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
