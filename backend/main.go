package main

import (
	"fmt"
	"note-sync/config"
	"note-sync/controllers"
	"note-sync/middleware"
	"note-sync/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化数据库
	models.InitDB()

	r := gin.Default()

	// 配置CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 认证相关路由
	auth := r.Group("/auth")
	{
		auth.GET("/github/login", controllers.HandleGithubLogin)
		auth.GET("/github/callback", controllers.HandleGithubCallback)
	}

	// API路由
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/articles", controllers.GetArticles)
		api.POST("/articles", controllers.CreateArticle)
		api.PUT("/articles/:id", controllers.UpdateArticle)
		api.DELETE("/articles/:id", controllers.DeleteArticle)
		api.GET("/articles/:id", controllers.GetArticle)

		api.GET("/questions", controllers.GetQuestions)
		api.GET("/questions/:id", controllers.GetQuestion)
		api.POST("/questions", controllers.CreateQuestion)
		api.PUT("/questions/:id", controllers.UpdateQuestion)
		api.DELETE("/questions/:id", controllers.DeleteQuestion)

		api.GET("/user", controllers.GetCurrentUser)
	}

	port := fmt.Sprintf(":%s", config.AppConfig.Server.Port)
	r.Run(port)
}
