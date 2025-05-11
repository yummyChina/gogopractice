package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;type:varchar(50)"`
	Password string `gorm:"not null;type:varchar(100)"`
	Email    string `gorm:"unique;not null;type:varchar(500)"`
}
