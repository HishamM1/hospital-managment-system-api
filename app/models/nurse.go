package models

type Nurse struct {
	// gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Patients []*Patient `gorm:"many2many:patient_nurses;"`
}
