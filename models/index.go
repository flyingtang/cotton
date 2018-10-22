package models

import (
	"github.com/jinzhu/gorm"
	"cotton/models/rbac"
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
	db.AutoMigrate(&Account{})

	// 用户组-n-> 用户
	db.Model(&rbac.UsernameGroup{}).Related(&Account{})
	db.Model(&rbac.UsernameGroup{}).Related(&rbac.Role{})
	db.Model(&rbac.Role{}).Related(&Account{})
	return db
}

