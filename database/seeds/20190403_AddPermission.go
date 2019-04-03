package seeds

import "github.com/RyotaNakajimakun/blog/models"

func AddPermissions() []interface{} {
	var permissions []interface{}

	permissions = append(permissions, models.NewPermission("view.development_menu.all", "開発中メニュー閲覧権限", "開発者専用"))
	permissions = append(permissions, models.NewPermission("edit.development_menu.all", "開発中メニュー編集権限", "開発者専用"))
	permissions = append(permissions, models.NewPermission("create.development_menu.all", "開発中メニュー作成権限", "開発者専用"))

	return permissions
}
