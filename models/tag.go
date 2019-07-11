package models

//Tag struct contains post tag info
type Tag struct {
	Model

	Title string `binding:"required" form:"title" gorm:"primary_key"`
	Posts []Post `gorm:"many2many:posts_tags;foreignkey:title"`
}
