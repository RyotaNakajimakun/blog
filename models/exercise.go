package models

type Exercise struct {
	TrainingID   int    `gorm:"not null"`
	TrainingName string `form:"training_name" gorm:"not null;default:''"`
	WeightID     int    `form:"weight_id"     gorm:"not null;default:0"`
	Rep          int    `form:"rep"           gorm:"not null;default:0"`
	Set          int    `form:"set"           gorm:"not null;default:0"`
	Interval     int    `form:"interval"      gorm:"not null;default:0"`
	Weight       Uint //※Uintはモデル名
}
