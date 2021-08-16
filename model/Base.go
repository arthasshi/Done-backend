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
	dsn := "vip:8DPRDTGZK5BphnX7@tcp(103.153.139.80:3306)/vip?charset=utf8mb4&parseTime=True&loc=Local"
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
