package admin

import (
	"github.com/RyotaNakajimakun/blog/controllers"
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"net/http"
)

func RoleIndex(c *gin.Context) {
	db := models.GetDB()
	var roles []models.Role
	db.Find(&roles)
	h := controllers.DefaultH(c)
	h["Title"] = "ロール一覧"
	h["Roles"] = roles

	c.HTML(http.StatusOK, "role/index", h)
}

func RoleNew(c *gin.Context) {
	h := controllers.DefaultH(c)
	h["Title"] = "権限の追加"
	c.HTML(http.StatusOK, "role/form", h)
}

func RoleCreate(c *gin.Context) {
	role := models.Role{}
	db := models.GetDB()

	if err := c.ShouldBind(&role); err != nil {
		session := sessions.Default(c)
		session.AddFlash(err.Error())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/admin/new_role")
		return
	}

	if err := db.Create(&role).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/role")
}

func RoleDetail(c *gin.Context) {
	var role models.Role
	var permissions []models.Permission
	type Permission struct {
		models.Permission
		Check bool
	}
	var HasPermissions []Permission

	db := models.GetDB()
	db.Where("name = ?", c.Param("name")).First(&role)
	db.Model(&role).Related(&role.HasPermission, "HasPermission")
	db.Find(&permissions)

	has := make([]bool, len(permissions))

	for i := range role.HasPermission {
		var perId int = role.HasPermission[i].ID - 1
		has[perId] = true
	}

	for i := range permissions {
		HasPermissions = append(HasPermissions, Permission{})
		HasPermissions[i].ID = permissions[i].ID
		HasPermissions[i].Name = permissions[i].Detail
		HasPermissions[i].DisplayName = permissions[i].DisplayName
		HasPermissions[i].Detail = permissions[i].Detail
		HasPermissions[i].Check = has[i]
	}

	h := controllers.DefaultH(c)
	h["Title"] = "権限詳細"
	h["HasPermissions"] = HasPermissions
	c.HTML(http.StatusOK, "role/detail", h)
}

func RoleEdit(c *gin.Context) {
	var permissions []models.Permission
	db := models.GetDB()
	db.Find(&permissions)

	h := controllers.DefaultH(c)
	h["Title"] = "権限編集"
	h["Permissions"] = permissions
	c.HTML(http.StatusOK, "role/edit", h)
}

func ChangePermission(c *gin.Context) {

}
