package controllers

import (
	"main/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

var (
	g          = galidator.New()
	customizer = g.Validator(doctorRequest{})
)

type doctorRequest struct {
	Name           string `json:"Name" binding:"required,min=3,max=255" required:"$field is required" min:"$field must be at least 3 characters" max:"$field must be at most 255 characters"`
	Email          string `json:"Email" binding:"required,email" required:"$field is required" email:"$field must be a valid email"`
	Number         string `json:"Number" binding:"required,min=3,max=255" required:"$field is required" min:"$field must be at least 3 characters" max:"$field must be at most 255 characters"`
	Specialization string `json:"Specialization" binding:"required,min=1,max=255" required:"$field is required" min:"$field must be at least 1 characters" max:"$field must be at most 255 characters"`
}

type doctorResponse struct {
	doctorRequest
	ID uint `json:"ID"`
}

func CreateDoctor(c *gin.Context) {
	var data doctorRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": customizer.DecryptErrors(err)})
		return
	}

	doctor := models.Doctor{}
	doctor.Name = data.Name
	doctor.Email = data.Email
	doctor.Number = data.Number
	doctor.Specialization = data.Specialization

	result := db.Create(&doctor)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response doctorResponse
	response.ID = doctor.ID
	response.Name = doctor.Name
	response.Email = doctor.Email
	response.Number = doctor.Number
	response.Specialization = doctor.Specialization

	c.JSON(200, response)
}

func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor

	result := db.Find(&doctors)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, doctors)
}

func GetDoctor(c *gin.Context) {
	var doctor models.Doctor

	result := db.First(&doctor, c.Param("id"))

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Doctor not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, doctor)
}

func UpdateDoctor(c *gin.Context) {
	var data doctorRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": customizer.DecryptErrors(err)})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	// check if doctor exists
	var doctor models.Doctor

	result := db.First(&doctor, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Doctor not found"})
		return
	}

	doctor.ID = uint(id)
	doctor.Name = data.Name
	result = db.Save(&doctor)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response doctorResponse
	response.ID = doctor.ID
	response.Name = doctor.Name
	response.Email = doctor.Email
	response.Number = doctor.Number
	response.Specialization = doctor.Specialization

	c.JSON(200, response)
}

func DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var doctor models.Doctor

	result := db.First(&doctor, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Doctor not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	result = db.Delete(&doctor)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Doctor deleted"})
}
