package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInvestor(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Investor

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	investor := models.Investor{Name: body.Name, UserID: user.(models.User).ID}
	result := initializers.DB.Create(&investor)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create investor",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Investor created successfully",
		"investor": investor,
	})
}

func GetInvestors(c *gin.Context) {
	user, _ := c.Get("user")
	var investors []models.Investor

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Find(&investors)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get investors",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"investors": investors,
	})
}
