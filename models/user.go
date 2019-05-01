package models

import (
	"golang.org/x/crypto/bcrypt"
)

//Login view model
type Login struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//Register view model
type Register struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//User type contains user info
type User struct {
	Model

	Email    string `form:"email" binding:"required" gorm:"unique"`
	Name     string `form:"name"`
	Password string `form:"password" binding:"required"`
	RoleID   uint64 `gorm:"default:4"`
	Role     Role
}

//BeforeSave gorm hook
func (u *User) BeforeSave() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hash)
	return
}

//WARNING:セキュリティ的にどうなのか検討が必要
func InitializeUser(uID interface{}) *User {
	db := GetDB()
	user := User{}
	db.First(&user, uID)
	db.Model(&user).Related(&user.Role, "Role")
	db.Model(&user.Role).Related(&user.Role.HasPermission, "HasPermission")
	return &user
}
