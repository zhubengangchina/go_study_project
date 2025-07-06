package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance float64
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_task?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取底层连接失败")
	}

	db.AutoMigrate(&Account{}, &Transaction{})

	// db.Create(&Account{Balance: 1000})
	// db.Create(&Account{Balance: 2000})
	//发起转账
	result := TransferMoney(db, 1, 2, 100)
	if result != nil {
		fmt.Println("转账失败：", result)
	} else {
		fmt.Println("转账成功")
	}

}

func TransferMoney(db *gorm.DB, fromID uint, toID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account

		//查询来源账号
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, fromID).Error; err != nil {
			return err
		}

		//检查余额
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}

		//查询目标账户
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount, toID).Error; err != nil {
			return err
		}

		//修改余额
		fromAccount.Balance -= amount
		toAccount.Balance += amount

		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}

		//写入交易记录
		trasations := Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}
		if err := tx.Create(&trasations).Error; err != nil {
			return err
		}

		return nil
	})
}
