package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	CECODE int //client request error code
	SECODE int // server error code
)

func init() {
	// dsn := "root:Vip_root123@tcp(45.76.54.104:3306)/vip?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:Vip_root123@tcp(185.251.249.143:3306)/vip?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:Vip_root123@tcp(0.0.0.0:3306)/vip?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&User{})
	}
}

type ResJsonType struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
