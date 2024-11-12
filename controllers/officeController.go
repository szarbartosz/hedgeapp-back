package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOffice(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Office

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	office := models.Office{
		Name: body.Name,
		Address: models.Address{
			UserID:  user.(models.User).ID,
			City:    body.Address.City,
			Street:  body.Address.Street,
			Number:  body.Address.Number,
			ZipCode: body.Address.ZipCode,
		}}

	result := initializers.DB.Create(&office)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create office",
		})
		return
	}

	c.JSON(http.StatusOK, office)
}

func UpdateOffice(c *gin.Context) {
	var body models.Office

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var office models.Office
	result := initializers.DB.Preload("Address").Where("id = ?", c.Param("id")).First(&office)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Office not found",
		})
		return
	}

	office.Name = body.Name

	office.Address.City = body.Address.City
	office.Address.Street = body.Address.Street
	office.Address.Number = body.Address.Number
	office.Address.ZipCode = body.Address.ZipCode

	tx := initializers.DB.Begin()

	if err := tx.Save(&office).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update office",
		})
		return
	}

	if err := tx.Save(&office.Address).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update address",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, office)
}

func GetOffices(c *gin.Context) {
	var offices []models.Office

	result := initializers.DB.Preload("Address").Find(&offices)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get offices",
		})
		return
	}

	c.JSON(http.StatusOK, offices)
}

func GetSingleOffice(c *gin.Context) {
	user, _ := c.Get("user")
	var office models.Office

	result := initializers.DB.Preload("Address").Preload("Locations", "user_id = ?", user.(models.User).ID).Preload("Locations.Status").First(&office, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get office",
		})
		return
	}

	c.JSON(http.StatusOK, office)
}

func DeleteOffice(c *gin.Context) {
	var office models.Office

	result := initializers.DB.First(&office, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Office not found",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&office)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete office: " + deleteResult.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Office deleted",
	})
}
