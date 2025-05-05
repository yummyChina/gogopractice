package main

import (
	"com.example.advancegorm/constant"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password  string `gorm:"type:varchar(100);not null"`
	Posts     []Post
	PostCount int `gorm:"default:0"`
}

type Post struct {
	gorm.Model
	Title         string `gorm:"type:varchar(200);not null"`
	Content       string `gorm:"type:text;not null"`
	UserID        uint
	User          User `gorm:"foreignkey:UserID"`
	Comments      []Comment
	CommentStatus string `gorm:"type:varchar(20);default:'无评论'"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint
	Post    Post `gorm:"foreignkey:PostID"`
}

func main() {
	db, err := gorm.Open(sqlite.Open(constant.DBPATH))
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})

	if err != nil {
		panic(err)
	}

	tables, _ := db.Migrator().GetTables()
	fmt.Println("现有数据表", tables)

	userID := 1
	user, err := getUserPostWithComments(db, uint(userID))
	if err != nil {
		fmt.Println("查询失败", err)
	} else {
		for _, post := range user.Posts {
			fmt.Printf("文章[%d] 《%s》有%d条 评论", post.ID, post.Title, len(post.Comments))
			for _, comment := range post.Comments {
				fmt.Printf("     %s\n", comment.Content)
			}
		}
	}
}

func getUserPostWithComments(db *gorm.DB, userID uint) (User, error) {
	var user User
	result := db.Preload("Posts.Comments").Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func getMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	result := db.Model(&Post{}).
		Select("post.*,COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post)
	return post, result.Error
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id=?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count+1")).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id= ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	status := "有评论"
	if count == 0 {
		status = "无评论"
	}

	return tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comment_status", status).Error

}
