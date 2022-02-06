package basic

import "testing"

//type Point struct { X, Y float64}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path1 []Point

func (path Path1) TranslateBy(offset Point, add bool) Path1 {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
	return path
}

func TestFuncValue(t *testing.T) {
	path := Path1{Point{1, 2}, Point{2, 3}}
	t.Log(path.TranslateBy(Point{1, 1}, true))
	t.Log(path.TranslateBy(Point{1, 1}, false))
}