package router

import (
	"gin-gorm-demo/handlers"
	"gin-gorm-demo/middleware"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LogInfoMiddleware())
	defaultGroup := r.Group("/api")
	defaultGroup.Use(middleware.JWTAuth())
	defaultGroup.POST("/posts", handlers.CreatePost)
	defaultGroup.PUT("/posts", handlers.UpdatePost)
	defaultGroup.DELETE("/posts/:id", handlers.DeletePost)
	defaultGroup.POST("/comments", handlers.AddComment)
	r.GET("/posts", handlers.GetAllPosts)
	r.GET("/posts/:id", handlers.GetPost)
	r.GET("/posts/:id/comments", handlers.GetCommentsByPost)

	noAuthGroup := r.Group("api")
	noAuthGroup.POST("/login", handlers.Login)
	noAuthGroup.POST("/register", handlers.Register)

	return r
}
