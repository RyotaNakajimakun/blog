package main

import (
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/RyotaNakajimakun/blog/system"
)

func CreateTableRolePermission() {
	InitDatabase()
	db := models.GetDB()
	db.AutoMigrate(&models.Role{}, &models.Permission{})

	permission := models.Permission{
		Name:        "test_permission",
		DisplayName: "test_name",
		Detail:      "test_detail",
	}

	role := models.Role{
		Name:        "development",
		DisplayName: "開発者用権限",
		Detail:      "全機能の使用許可も持つ",
	}
	db.NewRecord(permission)
	db.NewRecord(role)
}
func InitDatabase() {
	system.LoadConfig()
	models.SetDB(system.GetConnectionString())
}

func main() {
	CreateTableRolePermission()
}
