package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStatus(c *gin.Context) {
	var body models.Status

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	status := models.Status{Name: body.Name}
	result := initializers.DB.Create(&status)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func UpdateStatus(c *gin.Context) {
	var body models.Status

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result := initializers.DB.Where("id", c.Param("id")).Updates(&body)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status updated successfully",
	})
}

func GetStatuses(c *gin.Context) {
	var statuses []models.Status

	result := initializers.DB.Find(&statuses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get statuses",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statuses": statuses,
	})
}

func DeleteStatus(c *gin.Context) {
	var status models.Status

	result := initializers.DB.First(&status, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Status not found",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&status)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete investor: " + deleteResult.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Status deleted",
	})
}
