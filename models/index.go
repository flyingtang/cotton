package models

import (
	"github.com/jinzhu/gorm"
	"cotton/models/account"
)


const mysqlUrl = "root:root@/cotton?charset=utf8&parseTime=True&loc=Local"

var DB *gorm.DB

func NewMySQL() *gorm.DB{

	db, err:= gorm.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err.Error())
	}
	if DB != nil {
		DB.Close()
	}
	DB = db

	account.InitTable(db)

	return db
}

