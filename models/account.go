package models

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model

	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Nickname string	`gorm:"size:255"`
}



