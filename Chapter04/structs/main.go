package main

import "fmt"

// 定义结构体
type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

type point struct {
	x int
	y int
}

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	var dot1 dot
	dot2 := dot{}
	dot2.name = "A"
	// 使用提升的字段名
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	// 初始化
	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	dot4 := dot{}
	// 使用类型名
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

func compare() (bool, bool) {
	// 匿名结构体
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	point2 := struct {
		x int
		y int
	}{}
	point2.x = 10
	point2.y = 5

	point3 := point{10, 10}

	return point1 == point2, point1 == point3
}

func getUsers() []user {
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}

	u2 := user{
		age:  19,
		name: "Nick",
	}

	// 未初始化，默认为对应字段类型的零值
	var u3 user
	u3.name = "Sue"
	u3.age = 31
	u3.member = true
	u3.balance = 17.09

	return []user{u1, u2, u3}
}

func main() {
	users := getUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}

	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)

	// 组合结构体
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}
}
