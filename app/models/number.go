package models

type Number struct {
	// gorm.Model
	ID            uint `gorm:"primaryKey"`
	PatientID     uint
	PatientNumber string
	FamilyNumber  string
	Patient       *Patient
}
