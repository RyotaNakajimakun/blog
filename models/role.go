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

func HasPermission(name string) interface{} {
	type PermissionCheck struct {
		Permission
		Check bool
	}
	var HasPermissions []PermissionCheck
	var role Role
	var permissions []Permission

	db := GetDB()
	db.Where("name = ?", name).First(&role)
	db.Model(&role).Related(&role.HasPermission, "HasPermission")
	db.Find(&permissions)

	has := make([]bool, len(permissions))
	for i := range role.HasPermission {
		var permissionID int = role.HasPermission[i].ID - 1
		has[permissionID] = true
	}

	for i := range permissions {
		HasPermissions = append(HasPermissions, PermissionCheck{})
		HasPermissions[i].ID = permissions[i].ID
		HasPermissions[i].Name = permissions[i].Name
		HasPermissions[i].DisplayName = permissions[i].DisplayName
		HasPermissions[i].Detail = permissions[i].Detail
		HasPermissions[i].Check = has[i]
	}
	return HasPermissions
}
