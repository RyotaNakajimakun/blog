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
	if perId, ok := c.GetPostFormArray("ID"); ok{
		pp.Print(perId)
	}
	if Changed, ok := c.GetPostFormArray("hasChanged"); ok{
		pp.Print(Changed)
	}
	c.Redirect(http.StatusOK, "admin/role/index")
}
