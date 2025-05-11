package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null;type:varchar(50)" json:"title" binding:"required"`
	Content string `gorm:"not null" ;type:"text" json:"content" binding:"required"`
	UserID  uint   `json:"user_id" binding:"required"`
}
