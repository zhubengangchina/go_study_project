package main

import "fmt"

/*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("员工信息：")
	fmt.Println("姓名:", e.Name)
	fmt.Println("年龄:", e.Age)
	fmt.Println("员工编号:", e.EmployeeID)
}

func main() {
	e := Employee{
		Person:     Person{Name: "1241", Age: 1},
		EmployeeID: "121",
	}
	e.PrintInfo()
}
