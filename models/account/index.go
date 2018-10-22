package account


import (
	"github.com/jinzhu/gorm"
	"cotton/models"
	"fmt"
)

type Account struct {
	gorm.Model

	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Nickname string	`gorm:"size:255"`
}



func Find() *gorm.DB{

	var account Account
	var se = []string{"username, nickname"}
	db := models.DB.Select(se).Find(&account)
	return db
	fmt.Println(db, "123")
}
