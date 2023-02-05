package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDeveloper(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Developer

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	developer := models.Developer{Name: body.Name, UserID: user.(models.User).ID}
	result := initializers.DB.Create(&developer)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create developer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Developer created successfully",
		"developer": developer,
	})
}

func GetDevelopers(c *gin.Context) {
	user, _ := c.Get("user")
	var developers []models.Developer

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Find(&developers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get developers",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"developers": developers,
	})
}
