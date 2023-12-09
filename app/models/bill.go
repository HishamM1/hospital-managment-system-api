package models

import "time"

type Bill struct {
	ID        uint `gorm:"primaryKey"`
	PatientID uint
	Patient   Patient
	Amount    float64
	CreatedAt time.Time
}
