package models

type Patient struct {
	// gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	BirthDate   string
	Address     string
	Disease     string
	StartDate   string
	DoctorID    uint
	RoomID      uint
	TreatmentID uint
	Doctor      *Doctor
	Room        *Room
	Treatment   *Treatment
	Numbers     []*Number
	Nurses      []*Nurse `gorm:"many2many:patient_nurses;"`
}
