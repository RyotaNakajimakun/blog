package models

import "time"

type Training struct {
	Model

	Note         string     `form:"note" gorm:"not null;default:''"`
	TrainingDay  time.Time  `form:"training_day"`
	StartingTime time.Time  `form:"starting_time"`
	EndingTime   time.Time  `form:"ending_time"`
	Exercise     []Exercise `gorm:`
}
