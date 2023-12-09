package controllers

import (
	"main/app/models"
	"main/config"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type numberRequest struct {
	Number      string `json:"number" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"required,min=1,max=255"`
}

type patientRequest struct {
	Name        string          `json:"name" binding:"required"`
	BirthDate   string          `json:"birth_date" binding:"required"`
	Address     string          `json:"address" binding:"required"`
	Disease     string          `json:"disease" binding:"required"`
	StartDate   string          `json:"start_date" binding:"required"`
	DoctorID    uint            `json:"doctor_id" binding:"required"`
	RoomID      uint            `json:"room_id" binding:"required"`
	TreatmentID uint            `json:"treatment_id" binding:"required"`
	Numbers     []numberRequest `json:"numbers" binding:"required"`
	Nurses      []uint          `json:"nurses" binding:"required"`
}

type patientResponse struct {
	patientRequest
	ID uint `json:"id"`
}

var patient_val = g.Validator(patientRequest{})

func GetPatients(c *gin.Context) {
	var patients []models.Patient

	result := db.Preload("Doctor").Preload("Room").Find(&patients)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, patients)
}

func GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient models.Patient

	result := db.Preload("Doctor").Preload("Room").Preload("Treatment").Preload("Numbers").Preload("Nurses").First(&patient, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, patient)
}

func CreatePatient(c *gin.Context) {
	var data patientRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": patient_val.DecryptErrors(err)})
		return
	}

	// check if nurse exists
	var nurses []models.Nurse

	result := db.Find(&nurses, data.Nurses)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	if len(nurses) != len(data.Nurses) {
		c.JSON(400, gin.H{"error": "Nurse not found"})
		return
	}

	patient := models.Patient{}
	patient.Name = data.Name
	patient.BirthDate = data.BirthDate
	patient.Address = data.Address
	patient.Disease = data.Disease
	patient.StartDate = data.StartDate
	patient.DoctorID = data.DoctorID
	patient.RoomID = data.RoomID
	patient.TreatmentID = data.TreatmentID

	result = db.Create(&patient)

	for _, number := range data.Numbers {
		db.Model(&patient).Association("Numbers").Append(&models.Number{
			Number:      number.Number,
			Description: number.Description,
		})
	}

	for _, nurse := range nurses {
		db.Model(&patient).Association("Nurses").Append(&nurse)
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response patientResponse
	response.ID = patient.ID
	response.Name = patient.Name
	response.BirthDate = patient.BirthDate
	response.Address = patient.Address
	response.Disease = patient.Disease
	response.StartDate = data.StartDate
	response.DoctorID = patient.DoctorID
	response.RoomID = patient.RoomID
	response.TreatmentID = patient.TreatmentID
	response.Numbers = data.Numbers
	response.Nurses = data.Nurses

	c.JSON(200, response)
}

func UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var data patientRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": patient_val.DecryptErrors(err)})
		return
	}

	// check if nurse exists
	var nurses []models.Nurse

	result := db.Find(&nurses, data.Nurses)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	if len(nurses) != len(data.Nurses) {
		c.JSON(400, gin.H{"error": "Nurse not found"})
		return
	}

	// check if doctor exists
	var doctor models.Doctor

	result = db.First(&doctor, data.DoctorID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Doctor not found"})
		return
	}

	// check if room exists
	var room models.Room

	result = db.First(&room, data.RoomID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}

	// check if treatment exists

	var treatment models.Treatment

	result = db.First(&treatment, data.TreatmentID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Treatment not found"})
		return
	}

	var patient models.Patient

	result = db.Preload("Doctor").Preload("Room").Preload("Treatment").Preload("Numbers").Preload("Nurses").First(&patient, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	patient.Name = data.Name
	patient.BirthDate = data.BirthDate
	patient.Address = data.Address
	patient.Disease = data.Disease
	patient.StartDate = data.StartDate
	patient.DoctorID = data.DoctorID
	patient.RoomID = data.RoomID
	patient.TreatmentID = data.TreatmentID

	result = db.Save(&patient)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	// delete numbers and assign new numbers
	var numbers []models.Number
	db.Model(&patient).Association("Numbers").Find(&numbers)
	db.Model(&patient).Association("Numbers").Clear()

	for _, number := range numbers {
		db.Delete(&number)
	}

	for _, number := range data.Numbers {
		db.Model(&patient).Association("Numbers").Append(&models.Number{
			Number:      number.Number,
			Description: number.Description,
		})
	}

	// delete nurses and assign new nurses
	db.Model(&patient).Association("Nurses").Clear()

	for _, nurse := range nurses {
		db.Model(&patient).Association("Nurses").Append(&nurse)
	}

	c.JSON(200, patient)
}

func DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid doctor ID"})
		return
	}

	var patient models.Patient

	result := db.First(&patient, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	// delete numbers
	var numbers []models.Number
	db.Model(&patient).Association("Numbers").Find(&numbers)
	db.Model(&patient).Association("Numbers").Clear()

	for _, number := range numbers {
		db.Delete(&number)
	}

	// delete nurses
	db.Model(&patient).Association("Nurses").Clear()

	result = db.Delete(&patient)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Patient deleted successfully"})
}
