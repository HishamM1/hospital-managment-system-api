package models

import "time"

type Patient struct {
	// gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	BirthDate   string
	Address     string
	Disease     string
	StartDate   time.Time
	DoctorID    uint
	RoomID      uint
	TreatmentID uint
	Doctor      *Doctor
	Room        *Room
	Treatment   *Treatment
	Numbers     []*Number
	Nurses      []*Nurse `gorm:"many2many:patient_nurses;"`
}
