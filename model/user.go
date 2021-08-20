package model

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Birth    string `json:"birth"`
	Status   int    `json:"status"`   //会员状态
	Role     int    `json:"role"`     //角色包含：admin:0，门店管理员:1，会员:2，也就是说这些角色共用一个表
	Integral int    `json:"integral"` //会员积分
	ShopId   uint   `json:"shop_id"`
}

func RegistUser(u *User) (uint, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	u.Password = string(hash)
	result := db.Create(u)
	return u.ID, result.Error
}
func Login(u *User) (*User, error) {
	var user *User
	result := db.Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err == nil {
		user.Password = ""
		return user, nil
	} else {
		return nil, errors.New("用户名或密码错误")
	}
}

/*
	获取用户列表。靠status区分是哪个类型的用户
*/
func UserList(page, size, role int) ([]*User, int64, error) {
	var users []*User
	var total int64
	result := db.Table("users").Where("role = ?", role).Count(&total).Offset((page - 1) * size).Limit(size).Find(&users)
	return users, total, result.Error
}

func ChangeStatus(status int, id uint) (*User, error) {
	var user User
	db.First(&user, id)
	user.Status = status
	result := db.Model(&user).Update("status", status)
	return &user, result.Error
}
func ChangeUserInte(money float64, id uint) (*User, error) {
	var user *User
	result := db.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	if user.ShopId == 0 {
		return user, errors.New("会员没有所属门店，无法修改积分")
	}
	var shop *Shop
	result = db.First(&shop, user.ShopId)
	if result.Error != nil {
		return user, errors.New("会员没有所属门店，无法修改积分")
	}

	integral := user.Integral + int(money/shop.MoneyScale)
	result = db.Model(user).Update("integral", integral)
	return user, result.Error
}
