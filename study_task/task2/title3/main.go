package main

import (
	"fmt"
	"math"
)

/*
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {

	//面积
	Area() float64

	//周长
	Perimeter() float64
}

// 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// 园
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	//创建Rectangle 实力
	r := Rectangle{Width: 5, Height: 3}
	fmt.Println("矩形 Area:", r.Area())
	fmt.Println("矩形 Perimeter:", r.Perimeter())

	c := Circle{Radius: 4}
	fmt.Println("圆形 Area:", c.Area())
	fmt.Println("圆形 Perimeter:", c.Perimeter())
}
