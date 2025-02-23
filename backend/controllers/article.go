package controllers

import (
	"log"
	"net/http"
	"note-sync/models"
	"note-sync/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	// 从中间件获取用户ID
	userID, _ := c.Get("userID")

	var articles []models.Article
	if err := models.DB.Where("user_id = ?", userID).Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get articles"})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func CreateArticle(c *gin.Context) {
	userID, _ := c.Get("userID")
	githubToken, _ := c.Get("githubToken")

	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.UserID = userID.(uint)

	if err := models.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// 使用GitHub token同步到GitHub
	if err := services.SyncArticleToGithub(&article, githubToken.(string)); err != nil {
		log.Printf("Failed to sync to GitHub: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync to GitHub"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	// 从中间件获取用户ID
	userID, _ := c.Get("userID")

	id := c.Param("id")
	var article models.Article

	// 检查文章是否存在且属于当前用户
	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 更新文章
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	// 同步到GitHub
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if err := services.SyncArticleToGithub(&article, user.GithubToken); err != nil {
		log.Printf("Failed to sync to GitHub: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync to GitHub"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func DeleteArticle(c *gin.Context) {
	// 从中间件获取用户ID
	userID, _ := c.Get("userID")

	id := c.Param("id")
	var article models.Article

	// 检查文章是否存在且属于当前用户
	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 从GitHub删除
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if err := services.DeleteArticleFromGithub(&article, user.GithubToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete from GitHub"})
		return
	}

	// 从数据库删除
	if err := models.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}

func GetArticle(c *gin.Context) {
	// 从中间件获取用户ID
	userID, _ := c.Get("userID")

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	var article models.Article
	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&article).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}
