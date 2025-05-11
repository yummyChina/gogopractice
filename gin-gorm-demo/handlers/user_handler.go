package handlers

import "C"
import (
	"gin-gorm-demo/middleware"
	"gin-gorm-demo/model"
	"gin-gorm-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	result := model.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"})
	}
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User

	if err := model.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	if user.ID == 0 || !utils.VerifyPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	token, _ := middleware.GenerateJWT(user.ID, user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}
