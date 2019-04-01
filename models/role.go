package models


type Role struct {
	ID          int    `gorm:primary_key`
	Name        string `gorm:unique`
	DisplayName string `gorm:"not null"`
	Datail      string `gorm:"not null"`
	HasPermission []Permission `gorm:"many2many:permisson_role;"`
}

type Permission struct {
	ID          int    `gorm:"primary_key"`
	Name        string `gorm:"unique; not null"`
	DisplayName string `gorm:"not null"`
	Datail      string `gorm:"dafault:null"`
}
