package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type treatmentRequest struct {
	Type string `json:"type" binding:"required,min=1,max=255" required:"$field is required" min:"$field must be at least 1 characters" max:"$field must be at most 255 characters"`
}

type treatmentResponse struct {
	treatmentRequest
	ID uint `json:"id"`
}

var treatment_val = g.Validator(treatmentRequest{})

func CreateTreatment(c *gin.Context) {
	var data treatmentRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": treatment_val.DecryptErrors(err)})
		return
	}

	treatment := models.Treatment{}
	treatment.Type = data.Type

	result := db.Create(&treatment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response treatmentResponse
	response.ID = treatment.ID
	response.Type = treatment.Type

	c.JSON(200, response)
}

func GetTreatments(c *gin.Context) {
	var treatments []models.Treatment

	result := db.Find(&treatments)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, treatments)
}

func GetTreatment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var treatment models.Treatment

	result := db.First(&treatment, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Treatment not found"})
		return
	}

	c.JSON(200, treatment)
}

func UpdateTreatment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var data treatmentRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": treatment_val.DecryptErrors(err)})
		return
	}

	var treatment models.Treatment

	result := db.First(&treatment, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Treatment not found"})
		return
	}

	treatment.ID = uint(id)
	treatment.Type = data.Type

	result = db.Save(&treatment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response treatmentResponse
	response.ID = treatment.ID
	response.Type = treatment.Type

	c.JSON(200, response)
}

func DeleteTreatment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var treatment models.Treatment

	result := db.First(&treatment, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Treatment not found"})
		return
	}

	result = db.Delete(&treatment)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Treatment deleted successfully"})
}
