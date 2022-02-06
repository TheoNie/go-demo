package basic

import "testing"

type Employee struct {
	ID int
	Name, Address string
	Salary int
	ManagerId int
}

func getEmployeeByID(id int) *Employee {
	var a Employee
	a.ID = id
	a.Name = "name"
	a.Salary = 5000
	return &a
}

func TestStruct(t *testing.T) {
	t.Log(getEmployeeByID(2).Name)
	a := getEmployeeByID(2)
	a.Address = "address"
	t.Log(a)

	b := Employee{1, "name", "address", 5000, 0}
	t.Log(b)

	c := Employee{ID: 2, Address: "add"}
	t.Log(c)

	d := Employee{ID: 2, Address: "add"}
	t.Log(b == c)    //false
	t.Log(c == d)    //true
}

func TestAnonymousMem(t *testing.T) {
	type Point struct {
		x int
		y int
	}

	//存在非匿名成员
	type Circle struct {
		P Point
		Radius int
	}
	//存在匿名成员
	type Wheel struct {
		Point
		Spokes int
	}

	var circle Circle
	circle.P.x = 1
	circle.P.y = 1
	circle.Radius = 2

	var wheel Wheel
	wheel.x = 1
	wheel.y = 100
	wheel.Spokes = 300

	//通过匿名成员的方式可以很简单的组合一些小对象的功能
}