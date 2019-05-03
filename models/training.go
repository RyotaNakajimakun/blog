package models

import "time"

type Training struct {
	Model

	TrainingName string `form:"training_name" gorm:"not null;default:''"`
	WeightID     uint8  `form:"weight_id"     gorm:"not null;default:0"`
	Rep          uint16 `form:"rep"           gorm:"not null;default:0"`
	Set          uint8  `form:"set"           gorm:"not null;default:0"`
	Interval     uint16 `form:"interval"      gorm:"not null;default:0"`
	Weight       Weight
	Note         string `form:"note"          gorm:"not null;default:''"`

	TrainingDay  time.Time ``
	StartingTime time.Time ``
	EndingTime   time.Time ``
}
