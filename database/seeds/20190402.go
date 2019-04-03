package seeds

import (
	"fmt"
	"github.com/RyotaNakajimakun/blog/models"
)

func Seed() {
	var permissions []*models.Permission
	permission := NewPermission("admin","管理者用ロール","管理用権限を所持")
	permissions = append(permissions, permission)
	//permissions = append(permissions, NewPermission("admin","管理者用ロール","管理用権限を所持"))

	fmt.Print(permissions)
	//for i :=0; i < len(permissions); i++ {
	//	db.NewRecord(permissions[i])
	//}
}

func NewPermission(name string, display_name string, detail string) *models.Permission {
	permission := models.Permission{
		Name:        name,
		DisplayName: display_name,
		Detail:      detail,
	}
	fmt.Print(permission)
	return &permission
}


