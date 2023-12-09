package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type wardboyRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255" required:"$field is required" min:"$field must be at least 3 characters" max:"$field must be at most 255 characters"`
}

type wardboyResponse struct {
	wardboyRequest
	ID uint `json:"id"`
}

func CreateWardboy(c *gin.Context) {
	var data wardboyRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": customizer.DecryptErrors(err)})
		return
	}

	wardboy := models.Wardboy{}
	wardboy.Name = data.Name

	result := db.Create(&wardboy)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response wardboyResponse
	response.ID = wardboy.ID
	response.Name = wardboy.Name

	c.JSON(200, response)
}

func GetWardboys(c *gin.Context) {
	var wardboys []models.Wardboy

	result := db.Find(&wardboys)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, wardboys)
}

func GetWardboy(c *gin.Context) {
	var wardboy models.Wardboy

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	result := db.First(&wardboy, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Wardboy not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, wardboy)
}

func UpdateWardboy(c *gin.Context) {
	var data wardboyRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": customizer.DecryptErrors(err)})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var wardboy models.Wardboy

	result := db.First(&wardboy, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Wardboy not found"})
		return
	}

	wardboy.ID = uint(id)
	wardboy.Name = data.Name
	result = db.Save(&wardboy)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response wardboyResponse
	response.ID = wardboy.ID
	response.Name = wardboy.Name

	c.JSON(200, response)
}

func DeleteWardboy(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	var wardboy models.Wardboy

	result := db.First(&wardboy, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Wardboy not found"})
		return
	}

	result = db.Delete(&models.Wardboy{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Wardboy deleted successfully"})
}
