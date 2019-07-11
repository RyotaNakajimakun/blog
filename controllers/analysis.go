package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AnalysisIndex(c *gin.Context) {
	h := DefaultH(c)

	c.HTML(http.StatusOK, "analysis/index", h)
}

func AnalysisMyTraining(c *gin.Context) {
	h := DefaultH(c)

	c.HTML(http.StatusOK, "analysis/past_training", h)
}
