package controllers

import (
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"net/http"
)

func TrainingIndex(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "Training"

	c.HTML(http.StatusOK, "training/index", h)
}

func TrainingPost(c *gin.Context) {
	 h := DefaultH(c)
	 h["Title"] = "TrainingPost"

	 c.HTML(http.StatusOK, "training/post", h)
}

func TrainingSave(c *gin.Context) {
	training := models.Training{}
	c.ShouldBind(training)

	pp.Println(training)
	pp.Println(c.Request.Form)

	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/admin/training")
}
