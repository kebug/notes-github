package controllers

import (
	"context"
	"log"
	"net/http"

	"note-sync/config"
	"note-sync/models"
	"note-sync/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func HandleGithubLogin(c *gin.Context) {
	state := "random_state"

	// 创建带有代理客户端的 context
	_ = context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)

	// 使用带代理的 context 生成 URL
	url := config.GithubOauthConfig.AuthCodeURL(state)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGithubCallback(c *gin.Context) {
	code := c.Query("code")
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)

	token, err := config.GithubOauthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("Token exchange error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取用户信息并创建或更新用户
	user, err := models.CreateOrUpdateUser(token.AccessToken)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// 生成JWT token
	jwtToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Printf("Failed to generate JWT token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user":  user,
	})
}
