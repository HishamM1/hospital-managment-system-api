package models

type Number struct {
	// gorm.Model
	ID          uint `gorm:"primaryKey"`
	PatientID   uint
	Number      string
	Description string
	Patient     *Patient
}
