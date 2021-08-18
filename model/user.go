package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Birth    string `json:"birth"`
	Status   int    `json:"status"`
	Role     string `json:"role"` //
}

func RegistUser(u *User) (uint, error) {
	result := db.Create(u)
	return u.ID, result.Error
}

func UserList(page, size int) ([]*User, int64, error) {
	var users []*User
	var total int64
	result := db.Table("users").Count(&total).Offset((page - 1) * size).Limit(size).Find(&users)
	return users, total, result.Error
}
