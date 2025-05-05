package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"part1/constant"
)

type Student struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Grade string `gorm:"column:grade"`
}

type Account struct {
	ID      uint `gorm:"primary_key"`
	Balance float32
}

type Transaction struct {
	ID            uint    `gorm:"primary_key"`
	FromAccountID uint    `gorm:"column:from_account_id"`
	ToAccountID   uint    `gorm:"column:to_account_id"`
	Amount        float32 `gorm:"column:amount"`
}

func main() {
	db, err := gorm.Open(sqlite.Open(constant.DBPATH))
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	err = db.AutoMigrate(&Student{})

	if err != nil {
		panic(err)
	}
	//db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	//var students []Student
	//
	//db.Find(&students).Where("age > ?", 18)
	//
	//for _, student := range students {
	//	fmt.Println(student.Name, student.Age, student.Grade)
	//}

	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	err = db.Where("age < ?", 15).Delete(&Student{}).Error
	if err != nil {
		panic(err)
	}

}

func payMoney(aId uint, bId uint, db *gorm.DB, money float32) {
	var a Account
	var b Account
	err1 := db.Find(&a, "id = ?", aId)
	err2 := db.Find(&b, "id = ?", bId)

	if err1 != nil || err2 != nil {
		return
	}

	if a.Balance < money {
		fmt.Println("余额不足")
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&a).Update("balance", a.Balance-money).Error; err != nil {
			return err
		}

		if err := tx.Model(&b).Update("balance", b.Balance+money).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("交易失败")
	} else {
		fmt.Println("交易成功")
	}

}
