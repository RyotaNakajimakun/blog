package models

type Uint struct {
	ID   uint8  `gorm:"primary_key"`
	Name string `gorm:"unique"`
}
