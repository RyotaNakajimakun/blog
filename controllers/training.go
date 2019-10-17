package controllers

import (
	"github.com/RyotaNakajimakun/blog/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

//_@TODO Validationの処理を実装する必要がある
func TrainingSave(c *gin.Context) {
	var err error
	var trainingData models.TrainingForm
	formValues := c.Request.Form

	exercises := make([]models.Exercise, len(formValues["training_name"]))
	for cnt, exercise := range exercises {
		exercise.TrainingName = formValues["training_name"][cnt]
		exercise.Rep, err = strconv.Atoi(formValues["rep"][cnt])
		exercise.Set, err = strconv.Atoi(formValues["set"][cnt])
		exercise.Interval, err = strconv.Atoi(formValues["interval"][cnt])
		exercise.WeightID, err = strconv.Atoi(formValues["uint"][cnt])
		exercises[cnt] = exercise
	}

	db := models.GetDB()
	if err = c.ShouldBind(&trainingData); err != nil {
		session := sessions.Default(c)
		session.AddFlash(err.Error())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/admin/training")
		return
	}
	Data := models.SetTrainingData(trainingData, exercises)
	if err = db.Save(&Data).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/admin/training")
}
