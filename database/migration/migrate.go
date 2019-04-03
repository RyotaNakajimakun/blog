package migration

import (
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/RyotaNakajimakun/blog/system"
	"github.com/jinzhu/gorm"
)

func CreateTableRolePermission() {
	db := InitDatabase()
	db.AutoMigrate(&models.Role{}, &models.Permission{})
}
func InitDatabase() *gorm.DB{
	system.LoadConfig()
	models.SetDB(system.GetConnectionString())
	db := models.GetDB()
	return db
}
