package models

type Treatment struct {
	// gorm.Model
	ID   uint `gorm:"primaryKey"`
	Type string
}
