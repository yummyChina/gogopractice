package handlers

import (
	"gin-gorm-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePost(c *gin.Context) {
	userId := c.MustGet("userId").(uint)

	var post model.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = userId
	model.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{"message": "post create successfully"})
}

func UpdatePost(c *gin.Context) {

	var paramPost model.Post

	if err := c.ShouldBind(&paramPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	post, err := GetPostById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	loginId, _ := c.Get("userId")

	if loginId != paramPost.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := model.DB.Model(&post).Where("id = ?", id).Updates(post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post update successfully"})
}

func GetAllPosts(c *gin.Context) {
	var posts []model.Post
	model.DB.Preload("Comment").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	postId := c.MustGet("id").(uint)
	var post model.Post
	if err := model.DB.Where("id = ? ", postId).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)

}

func DeletePost(c *gin.Context) {
	postId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := GetPostById(uint(postId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	loginId, _ := c.Get("userId")
	if loginId != post.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err := model.DB.Model(&post).Delete(&model.Post{}, postId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post delete successfully"})
}

func GetPostById(id uint) (*model.Post, error) {
	var post model.Post
	err := model.DB.First(&post, id).Error
	return &post, err
}
