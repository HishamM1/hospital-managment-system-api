package models

type Wardboy struct {
	// gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}
