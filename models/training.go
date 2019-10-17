package models

import (
	"time"
)

type Training struct {
	Model

	Note         string     `form:"note" gorm:"not null;default:''"`
	TrainingDay  time.Time  `form:"training_day"`
	StartingTime time.Time  `form:"starting_time"`
	EndingTime   time.Time  `form:"ending_time"`
	Exercise     []Exercise `gorm:`
}

type TrainingForm struct {
	Note         string     `form:"note" gorm:"not null;default:''"`
	TrainingDay  string     `form:"training_day"`
	StartingTime string     `form:"starting_time"`
	EndingTime   string     `form:"ending_time"`
	Exercise     []Exercise `gorm:`
}

//_@OPTIMIZE レシーバーを利用して実装した方が無駄が少ない
func SetTrainingData(tf TrainingForm, exercises []Exercise) Training {
	var trainingData Training
	trainingData.Note = tf.Note
	trainingData.TrainingDay, _ = time.Parse("2006-01-02", tf.TrainingDay)
	trainingData.StartingTime, _ = time.Parse("15:04", tf.StartingTime)
	trainingData.EndingTime, _ = time.Parse("15:04", tf.EndingTime)

	trainingData.Exercise = exercises
	return trainingData
}
