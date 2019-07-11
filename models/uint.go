package models

type Uint struct {
	ID   int  `gorm:"primary_key"`
	Name string `gorm:"unique"`
}
