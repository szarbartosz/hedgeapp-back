package controllers

import (
	"main/initializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
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

	location := models.Location{
		UserID:            user.(models.User).ID,
		StatusID:          body.StatusID,
		InvestorID:        body.InvestorID,
		OfficeID:          body.OfficeID,
		Name:              body.Name,
		IssueDate:         body.IssueDate,
		InspectionDate:    body.InspectionDate,
		InspectionDone:    false,
		DecisionDate:      body.DecisionDate,
		DeforestationDate: body.DeforestationDate,
		DeforestationDone: false,
		PlantingDate:      body.PlantingDate,
		PlantingDone:      false,
		Application: models.Application{
			UserID:             user.(models.User).ID,
			Signature:          body.Application.Signature,
			IsCommercial:       body.Application.IsCommercial,
			DeforestationCause: body.Application.DeforestationCause,
			DeforestationDate:  body.Application.DeforestationDate,
			PlantingDate:       body.Application.PlantingDate,
			PlantingSite:       body.Application.PlantingSite,
			Species:            body.Application.Species,
		},
		Address: models.Address{
			UserID:  user.(models.User).ID,
			City:    body.Address.City,
			Street:  body.Address.Street,
			Number:  body.Address.Number,
			ZipCode: body.Address.ZipCode,
		}}

	result := initializers.DB.Create(&location)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create location",
		})
		return
	}

	c.JSON(http.StatusOK, location)
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
	result := initializers.DB.Preload("Application").Preload("Address").Where("user_id = ?", user.(models.User).ID).Where("id = ?", c.Param("id")).First(&location)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Location not found",
		})
		return
	}

	location.Name = body.Name
	location.IssueDate = body.IssueDate
	location.InspectionDate = body.InspectionDate
	location.InspectionDone = body.InspectionDone
	location.DecisionDate = body.DecisionDate
	location.DeforestationDate = body.DeforestationDate
	location.DeforestationDone = body.DeforestationDone
	location.PlantingDate = body.PlantingDate
	location.PlantingDone = body.PlantingDone
	location.UserID = user.(models.User).ID
	location.StatusID = body.StatusID
	location.InvestorID = body.InvestorID

	location.Application.Signature = body.Application.Signature
	location.Application.IsCommercial = body.Application.IsCommercial
	location.Application.DeforestationCause = body.Application.DeforestationCause
	location.Application.DeforestationDate = body.Application.DeforestationDate
	location.Application.PlantingDate = body.Application.PlantingDate
	location.Application.PlantingSite = body.Application.PlantingSite
	location.Application.Species = body.Application.Species

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

	if err := tx.Save(&location.Application).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update application",
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

	c.JSON(http.StatusOK, location)
}

func GetLocations(c *gin.Context) {
	user, _ := c.Get("user")
	var locations []models.Location

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Office.Address").Preload("Investor.Address").Preload(clause.Associations).Find(&locations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get locations",
		})
		return
	}

	c.JSON(http.StatusOK, locations)
}

func GetSingleLocation(c *gin.Context) {
	user, _ := c.Get("user")
	var location models.Location

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).Preload("Office.Address").Preload("Investor.Address").Preload(clause.Associations).First(&location, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get location",
		})
		return
	}

	c.JSON(http.StatusOK, location)
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

func AddNote(c *gin.Context) {
	user, _ := c.Get("user")
	var body models.Note
	var location models.Location

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result := initializers.DB.Where("user_id = ?", user.(models.User).ID).First(&location, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Location not found",
		})
		return
	}

	note := models.Note{
		LocationID: location.ID,
		Content:    body.Content,
	}
	result = initializers.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create note",
		})
		return
	}

	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	var body models.Note

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	result := initializers.DB.Where("id", c.Param("id")).Updates(&body)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update note",
		})
		return
	}

	var updatedNote models.Note
	initializers.DB.First(&updatedNote, c.Param("id"))

	c.JSON(http.StatusOK, updatedNote)
}

func DeleteNote(c *gin.Context) {
	var note models.Note

	result := initializers.DB.Where("id", c.Param("id")).First(&note)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Note not found",
		})
		return
	}

	deleteResult := initializers.DB.Delete(&note)

	if deleteResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete note: " + deleteResult.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted",
	})
}
