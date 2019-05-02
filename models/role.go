package models

type Role struct {
	ID            int          `gorm:"primary_key"`
	Name          string       `form:"name" binding:"required" gorm:"unique"`
	DisplayName   string       `form:"display_name" binding:"required" gorm:"not null"`
	Detail        string       `form:"detail" binding:"required" gorm:"not null"`
	HasPermission []Permission `gorm:"many2many:role_permission;"`
}

func NewRole(name string, displayName string, detail string) Role {
	role := Role{
		Name:        name,
		DisplayName: displayName,
		Detail:      detail,
	}
	return role
}
