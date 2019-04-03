package migration

import (
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/RyotaNakajimakun/blog/system"
)

func CreateTableRolePermission() {
	InitDatabase()
	db := models.GetDB()
	db.AutoMigrate(&models.Role{}, &models.Permission{})
}
func InitDatabase() {
	system.LoadConfig()
	models.SetDB(system.GetConnectionString())
}
