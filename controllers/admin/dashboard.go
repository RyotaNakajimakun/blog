package admin

import (
	"github.com/RyotaNakajimakun/blog/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AdminGet handles GET /admin/ route
func Show(c *gin.Context) {
	h := controllers.DefaultH(c)
	h["Title"] = "Admin dashboard"
	c.HTML(http.StatusOK, "admin/dashboard/show", h)
}
