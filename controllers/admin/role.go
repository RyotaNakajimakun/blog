package admin

import (
	"github.com/RyotaNakajimakun/blog/controllers"
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	HasPermissions := models.HasPermission(c.Param("name"))

	h := controllers.DefaultH(c)
	h["Title"] = "権限詳細"
	h["HasPermissions"] = HasPermissions
	c.HTML(http.StatusOK, "role/detail", h)
}

func RoleEdit(c *gin.Context) {
	HasPermissions := models.HasPermission(c.Param("name"))

	h := controllers.DefaultH(c)
	h["Title"] = "権限編集"
	h["HasPermissions"] = HasPermissions
	c.HTML(http.StatusOK, "role/edit", h)
}

func ChangePermission(c *gin.Context) {
	db := models.GetDB()
	role_name := c.Param("name")

	Role := models.Role{}
	db.Where("name = ?", role_name).First(&Role)

	if perId, ok := c.GetPostFormArray("permissionID"); ok {
		db.Model(&Role).Association("HasPermission").Clear()
		for _, ID := range perId {
			Permission := models.Permission{}
			db.First(&Permission, ID)
			db.Model(&Role).Association("HasPermission").Append(&Permission)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/admin/role")
}
