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

	investor := initializers.DB.Create(&models.Investor{
		Name:   body.Name,
		UserID: user.(models.User).ID,
		Address: models.Address{
			UserID:  user.(models.User).ID,
			City:    body.Address.City,
			Street:  body.Address.Street,
			Number:  body.Address.Number,
			ZipCode: body.Address.ZipCode,
		}})

	if investor.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create investor",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"investor": investor,
	})
}

func UpdateInvestor(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Investor

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var investor models.Investor
	result := initializers.DB.Preload("Address").Where("user_id = ?", user.(models.User).ID).Where("id = ?", c.Param("id")).First(&investor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Investor not found",
		})
		return
	}

	investor.Name = body.Name

	investor.Address.City = body.Address.City
	investor.Address.Street = body.Address.Street
	investor.Address.Number = body.Address.Number
	investor.Address.ZipCode = body.Address.ZipCode

	tx := initializers.DB.Begin()

	if err := tx.Save(&investor).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update investor",
		})
		return
	}

	if err := tx.Save(&investor.Address).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update address",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"investor": investor,
	})
}

func GetInvestors(c *gin.Context) {
	user, _ := c.Get("user")
	var investors []models.Investor

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Address").Preload("Locations").Find(&investors)

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

func GetSingleInvestor(c *gin.Context) {
	user, _ := c.Get("user")
	var investor models.Investor

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).First(&investor, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get investor",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"investor": investor,
	})
}

func DeleteInvestor(c *gin.Context) {
	user, _ := c.Get("user")
	var investor models.Investor

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).First(&investor, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Investor not found",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&investor)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete investor: " + deleteResult.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Investor deleted",
	})
}
