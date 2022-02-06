package basic

import (
	"math"
	"testing"
)

type Point struct {
	X, Y float64
}
type Path []Point

//传统方法
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//类型方法,接收器是非指针类型
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//类型方法,接收器是指针类型,区别是非指针类型对值的传递都是拷贝模式，同时不可对其值进行变更
//一般情况下，一个类型如果有任意一个方法是指针类型，那所有方法都要是指针类型
func (p *Point) Distance2(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//方法可以属于任何类型
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func TestMemFunc(t *testing.T)  {
	p := Point{1, 2}
	q := Point{2, 3}
	path := Path{p, q}
	t.Log(Distance(p, q))
	t.Log(p.Distance(q))
	t.Log(path.Distance())
	t.Log(p.Distance2(q)) //调用指针类型的方法时，也可以传原始值，会被自动隐式转换为指针
}
