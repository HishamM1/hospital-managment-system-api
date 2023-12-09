package models

type Room struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey"`
	Number    string
	Type      string
	WardboyID uint
	Wardboy   *Wardboy
}
