package models

type Doctor struct {
	// gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}
