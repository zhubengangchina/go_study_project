package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	db.AutoMigrate(&Student{})

	//插入学生信息
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	if err := db.Create(&student).Error; err != nil {
		fmt.Println("插入失败:", err)
	} else {
		fmt.Println("插入成功:", student)
	}

	//查询符合条件的学生
	var students []Student
	if err := db.Where("age > ?", 15).Find(&students).Error; err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Println("符合条件的学生：", students)
	}

	//修改学生信息
	if err := db.Model(&Student{}).
		Where("name = ?", "张三").
		Update("grade", "四年级").
		Error; err != nil {
		fmt.Println("更新失败:", err)
	} else {
		fmt.Println("更新成功")
	}

	//删除
	if err := db.Where("age < ?", 15).Delete(&Student{}).Error; err != nil {
		fmt.Println("删除失败:", err)
	} else {
		fmt.Println("删除成功")
	}

}
