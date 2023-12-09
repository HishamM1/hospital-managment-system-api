package controllers

import (
	"main/app/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type billRequest struct {
	PatientID uint      `json:"patient_id" binding:"required"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type billResponse struct {
	billRequest
	ID uint `json:"id"`
}

var bill_val = g.Validator(billRequest{})

func GetBills(c *gin.Context) {

	var bills []models.Bill

	result := db.Preload("Patient").Find(&bills)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, bills)
}

func GetBill(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var bill models.Bill
	result := db.Preload("Patient").First(&bill, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Bill not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, bill)
}

func CreateBill(c *gin.Context) {
	// count amount based on number of days patient stayed

	var data billRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": bill_val.DecryptErrors(err)})
		return
	}

	bill := models.Bill{}

	// check if patient exists
	var patient models.Patient
	result := db.First(&patient, data.PatientID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	bill.PatientID = data.PatientID

	currentDate := time.Now()
	startDate, _ := time.Parse("2006-01-02", patient.StartDate)
	difference := currentDate.Sub(startDate)

	// convert difference to days
	days := difference.Hours() / 24

	// multiply days by 1000 to get amount
	amount := days * 1000

	bill.Amount = amount

	result = db.Create(&bill)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response billResponse
	response.ID = bill.ID
	response.PatientID = bill.PatientID
	response.Amount = bill.Amount
	response.CreatedAt = bill.CreatedAt

	c.JSON(200, response)
}

func UpdateBill(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var data billRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": bill_val.DecryptErrors(err)})
		return
	}

	var bill models.Bill

	result := db.First(&bill, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Bill not found"})
		return
	}

	// check if patient exists
	var patient models.Patient
	result = db.First(&patient, data.PatientID)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	bill.ID = uint(id)
	bill.PatientID = data.PatientID
	bill.Amount = data.Amount

	result = db.Save(&bill)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	var response billResponse
	response.ID = bill.ID
	response.PatientID = bill.PatientID
	response.Amount = bill.Amount
	response.CreatedAt = bill.CreatedAt

	c.JSON(200, response)
}

func DeleteBill(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	result := db.Delete(&models.Bill{}, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Bill not found"})
		return
	}

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{"message": "Bill deleted successfully"})
}
