package models

type Weight struct {
	ID         uint8  `gorm:"primary_key"`
	WeightUnit string `gorm:"unique"`
}
