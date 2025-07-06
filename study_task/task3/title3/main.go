package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}

	//查询所有技术部的员工
	var employees []Employee
	err = db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatalln("查询技术部员工失败:", err)
	}
	fmt.Println("技术部员工:")
	for _, e := range employees {
		fmt.Printf("%+v\n", e)
	}

	//查询工资最高的员工
	var topEmployee Employee
	err = db.Get(&topEmployee, "select * from employees order by salary desc limit 1")
	if err != nil {
		log.Fatalln("查询最高薪资员工失败:", err)
	}
	fmt.Println("最高薪资员工:")
	fmt.Printf("%+v\n", topEmployee)

}
