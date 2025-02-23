package controllers

import (
	"log"
	"net/http"
	"note-sync/models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	// 从中间件获取用户ID
	userID, exists := c.Get("userID")
	if !exists {
		log.Printf("No userID in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 查询用户信息
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		log.Printf("Failed to get user %v: %v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	// 返回用户信息（不包含敏感字段）
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"githubId":  user.GithubID,
		"avatarUrl": user.AvatarURL,
	})
}
