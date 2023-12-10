package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NumberRequest struct {
	PatientID   uint   `json:"PatiendID" binding:"required"`
	Number      string `json:"Number" binding:"required,min=1,max=255"`
	Description string `json:"Description" binding:"required,min=1,max=255"`
}

type NumberResponse struct {
	NumberRequest
	ID uint `json:"id"`
}

var number_val = g.Validator(NumberRequest{})

func GetNumbers(c *gin.Context) {
	var numbers []models.Number

	result := db.Preload("Patient").Find(&numbers)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, numbers)
}

func GetNumber(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var number models.Number
	result := db.Preload("Patient").First(&number, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Number not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, number)
}

func CreateNumber(c *gin.Context) {
	var data NumberRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": number_val.DecryptErrors(err)})
		return
	}

	// check if patient exists
	var patient models.Patient
	result := db.First(&patient, data.PatientID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	number := models.Number{}
	number.PatientID = data.PatientID
	number.Number = data.Number
	number.Description = data.Description

	result = db.Create(&number)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response NumberResponse
	response.ID = number.ID
	response.PatientID = number.PatientID
	response.Number = number.Number
	response.Description = number.Description

	c.JSON(200, response)
}

func UpdateNumber(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var data NumberRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": number_val.DecryptErrors(err)})
		return
	}

	// check if patient exists
	var patient models.Patient
	result := db.First(&patient, data.PatientID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	var number models.Number
	result = db.First(&number, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Number not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	number.PatientID = data.PatientID
	number.Number = data.Number
	number.Description = data.Description

	result = db.Save(&number)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response NumberResponse
	response.ID = number.ID
	response.PatientID = number.PatientID
	response.Number = number.Number
	response.Description = number.Description

	c.JSON(200, response)
}

func DeleteNumber(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	number := models.Number{}
	result := db.Delete(&number, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Number not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Number deleted"})
}
