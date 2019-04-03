package models

type Role struct {
	ID            int          `gorm:"primary_key"`
	Name          string       `form:"name" binding:"required" gorm:"unique"`
	DisplayName   string       `form:"display_name" binding:"required" gorm:"not null"`
	Detail        string       `form:"detail" binding:"required" gorm:"not null"`
	HasPermission []Permission `gorm:"many2many:role_permission;"`
}

type Permission struct {
	ID          int    `form:"id" gorm:"primary_key"`
	Name        string `form:"name" gorm:"unique; not null"`
	DisplayName string `form:"display_name" gorm:"not null"`
	Detail      string `form:"detail" gorm:"default:null"`
}

func NewRole(name string, displayName string, detail string) Role {
	role := Role{
		Name:        name,
		DisplayName: displayName,
		Detail:      detail,
	}
	return role
}

func NewPermission(name string, displayName string, detail string) Permission {
	permission := Permission{
		Name:        name,
		DisplayName: displayName,
		Detail:      detail,
	}
	return permission
}

