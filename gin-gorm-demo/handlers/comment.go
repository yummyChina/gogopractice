package handlers

import (
	"gin-gorm-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	var commentReq model.Comment
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := c.Get("userId")

	comment := model.Comment{
		Content: commentReq.Content,
		PostID:  commentReq.PostID,
		UserID:  userId.(uint),
	}
	if err := model.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

func GetCommentsByPost(c *gin.Context) {
	postId, _ := strconv.ParseUint(c.Param("postId"), 10, 64)
	var comments []model.Comment
	if err := model.DB.Where("post_id= ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
