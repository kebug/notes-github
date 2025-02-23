package controllers

import (
	"log"
	"net/http"

	"note-sync/models"
	"note-sync/services"

	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	// 从中间件获取用户ID
	userID, _ := c.Get("userID")

	var questions []models.Question
	if err := models.DB.Where("user_id = ?", userID).Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get questions"})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func GetQuestion(c *gin.Context) {
	var question models.Question
	user := c.MustGet("user").(models.User)

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), user.ID).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "问题不存在"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func CreateQuestion(c *gin.Context) {
	userID, _ := c.Get("userID")
	githubToken, _ := c.Get("githubToken")

	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.UserID = userID.(uint)

	if err := models.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	// 同步到GitHub
	if err := services.SyncQuestionToGithub(&question, githubToken.(string)); err != nil {
		log.Printf("Failed to sync to GitHub: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync to GitHub"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func UpdateQuestion(c *gin.Context) {
	userID, _ := c.Get("userID")
	githubToken, _ := c.Get("githubToken")

	id := c.Param("id")
	var question models.Question

	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	// 同步到GitHub
	if err := services.UpdateQuestionInGithub(&question, githubToken.(string)); err != nil {
		log.Printf("Failed to sync to GitHub: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync to GitHub"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func DeleteQuestion(c *gin.Context) {
	userID, _ := c.Get("userID")
	githubToken, _ := c.Get("githubToken")

	id := c.Param("id")
	var question models.Question

	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&question).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// 从GitHub删除
	if err := services.DeleteQuestionFromGithub(&question, githubToken.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete from GitHub"})
		return
	}

	if err := models.DB.Delete(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
