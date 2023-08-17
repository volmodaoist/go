package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// 下面这种初始化方式非常类似于多层嵌套字典
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}

	// 下面这种初始化方式的可读性会更好一点
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Println(w1)
	fmt.Println(w2)

	// 结构体能够直接比较
	if w1 == w2 {
		fmt.Printf("&w1 = %p\n&w2 = %p\n", &w1, &w2)
		fmt.Println("They have the same attributes!")
	}

	// 匿名成员可以通过结构体类型名称访问成员,但是这种方法比较臃肿
	if w1.Circle.Point.X == w2.Circle.Point.X && w1.Circle.Point.Y == w2.Circle.Point.Y {
		fmt.Println("They have the same coordinate!")
	}

	// 也可以直接使用成员运算符访问内部成员, 这个特性在处理多层嵌套结构时非常优雅
	if w1.X == w2.X && w1.Y == w2.Y {
		fmt.Println("They have the same coordinate!")
	}

}
