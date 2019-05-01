package system

import (
	"github.com/RyotaNakajimakun/blog/models"
)

func authPermission(user models.User, authPermission string,) bool {
	if ok := _UserHasPermission(user, authPermission); ok {
		return true
	}
	return false
}

func _UserHasPermission(u models.User, authPermissionName string) bool {
	for _, HasPermission := range u.Role.HasPermission {
		if HasPermission.Name == authPermissionName {
			return true
		}
	}
	return false
}
