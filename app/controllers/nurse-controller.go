package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NurseRequest struct {
	Name string `json:"Name" binding:"required,min=1,max=255" required:"$field is required" min:"$field must be at least 1 characters" max:"$field must be at most 255 characters"`
}

type NurseResponse struct {
	NurseRequest
	ID   uint   `json:"ID"`
	Name string `json:"Name"`
}

var nurse_val = g.Validator(NurseRequest{})

func CreateNurse(c *gin.Context) {
	var data NurseRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": nurse_val.DecryptErrors(err)})
		return
	}

	nurse := models.Nurse{}
	nurse.Name = data.Name

	result := db.Create(&nurse)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response NurseResponse
	response.ID = nurse.ID
	response.Name = nurse.Name

	c.JSON(200, response)
}

func GetNurses(c *gin.Context) {
	var nurses []models.Nurse

	result := db.Find(&nurses)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, nurses)
}

func GetNurse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var nurse models.Nurse

	result := db.First(&nurse, id)

	if result.RowsAffected == 0 {
		c.JSON(400, gin.H{"error": "Nurse not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, nurse)
}

func UpdateNurse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var data NurseRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": nurse_val.DecryptErrors(err)})
		return
	}

	var nurse models.Nurse

	result := db.First(&nurse, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Nurse not found"})
		return
	}

	nurse.ID = uint(id)
	nurse.Name = data.Name

	result = db.Save(&nurse)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response NurseResponse
	response.ID = nurse.ID
	response.Name = nurse.Name

	c.JSON(200, response)
}

func DeleteNurse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var nurse models.Nurse

	result := db.First(&nurse, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Nurse not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	result = db.Delete(&nurse)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Nurse deleted successfully"})
}
