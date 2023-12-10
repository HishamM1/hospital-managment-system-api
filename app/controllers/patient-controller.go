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
	PatientNumber string `json:"PatientNumber" binding:"required,min=1,max=255"`
	FamilyNumber  string `json:"FamilyNumber" binding:"required,min=1,max=255"`
}

type patientRequest struct {
	Name        string        `json:"Name" binding:"required"`
	BirthDate   string        `json:"BirthDate" binding:"required"`
	Address     string        `json:"Address" binding:"required"`
	Disease     string        `json:"Disease" binding:"required"`
	StartDate   string        `json:"StartDate" binding:"required"`
	DoctorID    uint          `json:"DoctorID" binding:"required"`
	RoomID      uint          `json:"RoomID" binding:"required"`
	TreatmentID uint          `json:"TreatmentID" binding:"required"`
	Numbers     numberRequest `json:"Numbers" binding:"required"`
	Nurses      []uint        `json:"Nurses" binding:"required"`
}

type patientResponse struct {
	patientRequest
	ID        uint             `json:"ID"`
	Numbers   models.Number    `json:"Numbers"`
	Nurses    []models.Nurse   `json:"Nurses"`
	Doctor    models.Doctor    `json:"Doctor"`
	Room      models.Room      `json:"Room"`
	Treatment models.Treatment `json:"Treatment"`
}

var patient_val = g.Validator(patientRequest{})

func GetPatients(c *gin.Context) {
	var patients []models.Patient

	result := db.Preload("Doctor").Preload("Room").Preload("Treatment").Preload("Numbers").Preload("Nurses").Find(&patients)

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

	var response patientResponse
	response.ID = patient.ID
	response.Name = patient.Name
	response.BirthDate = patient.BirthDate
	response.Address = patient.Address
	response.Disease = patient.Disease
	response.StartDate = patient.StartDate
	response.Doctor = *patient.Doctor
	response.Room = *patient.Room
	response.Treatment = *patient.Treatment
	response.Numbers = models.Number{}
	response.Nurses = []models.Nurse{}

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

	db.Model(&patient).Association("Numbers").Append(&models.Number{
		PatientNumber: data.Numbers.PatientNumber,
		FamilyNumber:  data.Numbers.FamilyNumber,
	})

	for _, nurse := range nurses {
		db.Model(&patient).Association("Nurses").Append(&nurse)
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, patient)
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

	db.Model(&patient).Association("Numbers").Append(&models.Number{
		PatientNumber: data.Numbers.PatientNumber,
		FamilyNumber:  data.Numbers.FamilyNumber,
	})

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
