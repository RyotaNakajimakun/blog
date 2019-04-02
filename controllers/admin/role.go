package admin

import (
	"fmt"
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
	db := models.GetDB()
	role := models.Role{}
	var permissions []models.Permission
	db.Where("name = ?", c.Param("name")).First(&role)

	db.Model(&role).Related(&permissions, "HasPermission")

	fmt.Print(permissions)
	fmt.Print(role)

	h := controllers.DefaultH(c)
	c.HTML(http.StatusOK, "role/detail", h)
}

func RoleEdit(c *gin.Context) {
	h := controllers.DefaultH(c)
	c.HTML(http.StatusOK, "role/edit", h)
}
