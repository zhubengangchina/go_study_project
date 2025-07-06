package main

import (
	"errors"
	"fmt"
	"go_study_project/custom1/gorm_1/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	db.AutoMigrate(&models.User{})

	//创建用户A 和B
	userA := models.User{Name: "zhoudan55", Email: "ali55@example.com", Age: 25, Balance: 1000}
	userB := models.User{Name: "zhoudan66", Email: "ali66@example.com", Age: 25, Balance: 500}

	//db.Create([]models.User{userA, userB})
	db.Create(&userA)
	db.Create(&userB)

	err1 := TransferMoney(db, userA.ID, userB.ID, 200)
	if err1 != nil {
		fmt.Println("转账失败：", err)
	} else {
		fmt.Println("转账成功")
	}

	var updateA, updateB models.User
	db.First(&updateA, userA.ID)
	db.First(&updateB, userB.ID)
	fmt.Printf("用户A余额：%.2f\n", updateA.Balance)
	fmt.Printf("用户B余额：%.2f\n", updateB.Balance)

	RetryTransaction(db, 3, func(tx *gorm.DB) error {
		return TransferMoney(db, userA.ID, userB.ID, 200)
	})

	db.First(&updateA, userA.ID)
	db.First(&updateB, userB.ID)
	fmt.Printf("用户A余额：%.2f\n", updateA.Balance)
	fmt.Printf("用户B余额：%.2f\n", updateB.Balance)
}

func TransferMoney(db *gorm.DB, fromID uint, toID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		//查询转出用户的余额
		var fromUser models.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", fromID).First(&fromUser).Error; err != nil {
			return err
		}
		//检查余额是否狗
		if fromUser.Balance < amount {
			return errors.New("余额不足")
		}

		//扣除转出方余额
		if err := tx.Model(&models.User{}).
			Where("id = ?", fromID).
			Update("balance", gorm.Expr("balance - ?", amount)).
			Error; err != nil {
			return err
		}

		//增加接受余额
		if err := tx.Model(&models.User{}).
			Where("id = ?", toID).
			Update("balance", gorm.Expr("balance + ?", amount)).
			Error; err != nil {
			return err
		}
		return nil
	})
}

func RetryTransaction(db *gorm.DB, maxRetries int, txFunc func(tx *gorm.DB) error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err := db.Transaction(func(tx *gorm.DB) error {
			return txFunc(tx)
		})
		if err == nil {
			return nil
		}
		log.Printf("事务失败（第 %d 次）：%v\n", i+1, err)
		time.Sleep(100 * time.Millisecond) // 简单退避策略
	}
	return fmt.Errorf("事务多次重试失败：%w", err)
}
