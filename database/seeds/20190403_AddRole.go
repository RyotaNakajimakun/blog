package seeds

import (
	"github.com/RyotaNakajimakun/blog/models"
)

func AddRole() []models.Role {
	var roles []models.Role

	roles = append(roles, models.NewRole("development", "開発者用ロール", "全ての権限を所持"))
	roles = append(roles, models.NewRole("admin", "管理者用ロール", "管理用権限を所持"))
	roles = append(roles, models.NewRole("customer", "一般ユーザー用ロール", "一般公開権限を所持"))

	return roles
}
