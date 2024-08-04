package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Location

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	location := initializers.DB.Create(&models.Location{
		Name:              body.Name,
		IssueDate:         body.IssueDate,
		InspectionDate:    body.InspectionDate,
		DeforestationDate: body.DeforestationDate,
		DeforestationDone: false,
		PlantingDate:      body.PlantingDate,
		PlantingDone:      false,
		UserID:            user.(models.User).ID,
		StatusID:          body.StatusID,
		InvestorID:        body.InvestorID,
		Address: models.Address{
			UserID:  user.(models.User).ID,
			City:    body.Address.City,
			Street:  body.Address.Street,
			Number:  body.Address.Number,
			ZipCode: body.Address.ZipCode,
		}})

	if location.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create location",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location": location,
	})
}

func UpdateLocation(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Location

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var location models.Location
	result := initializers.DB.Preload("Address").Where("user_id = ?", user.(models.User).ID).Where("id = ?", c.Param("id")).First(&location)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Location not found",
		})
		return
	}

	location.Name = body.Name
	location.IssueDate = body.IssueDate
	location.InspectionDate = body.InspectionDate
	location.DeforestationDate = body.DeforestationDate
	location.DeforestationDone = body.DeforestationDone
	location.PlantingDate = body.PlantingDate
	location.PlantingDone = body.PlantingDone
	location.UserID = user.(models.User).ID
	location.StatusID = body.StatusID
	location.InvestorID = body.InvestorID

	location.Address.City = body.Address.City
	location.Address.Street = body.Address.Street
	location.Address.Number = body.Address.Number
	location.Address.ZipCode = body.Address.ZipCode

	tx := initializers.DB.Begin()

	if err := tx.Save(&location).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update location",
		})
		return
	}

	if err := tx.Save(&location.Address).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update address",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"location": location,
	})
}

func GetLocations(c *gin.Context) {
	user, _ := c.Get("user")
	var locations []models.Location

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Address").Find(&locations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get locations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"locations": locations,
	})
}

func GetSingleLocation(c *gin.Context) {
	user, _ := c.Get("user")
	var location models.Location

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Address").First(&location, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get location",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location": location,
	})
}

func DeleteLocation(c *gin.Context) {
	user, _ := c.Get("user")
	var location models.Location

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).First(&location, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Location not found",
		})
		return
	}

	initializers.DB.Delete(&location)

	if initializers.DB.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete location",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Location deleted",
	})
}
