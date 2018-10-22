package account


import (
	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
)

type Account struct {
	gorm.Model

	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Nickname string	`gorm:"size:255"`
	Enabled bool `gorm:"type:BOOL;default: true"`
}

var DB *gorm.DB





func Find(skip int64, where *string, arg ...string) (accounts []Account, total int64){

	var se = []string{"username, nickname, id, created_at, updated_at, deleted_at, enabled"}

	if where == nil {
		if res := DB.Limit(10).Offset(skip).Select(se).Find(&accounts).Count(&total);res.Error != nil {
			logrus.WithFields(logrus.Fields{"model": "Account Model", "action": "DB.Select(se).Find(&accounts)"}).Error(res.Error.Error())
		}
		return
	}
	if res := DB.Limit(10).Select(se).Offset(skip).Where(where, arg).Find(&accounts).Count(&total); res.Error != nil {
		logrus.WithFields(logrus.Fields{"model": "Account Model", "action": "DB.Select(se).Find(&accounts)"}).Error(res.Error.Error())
	}
	return
}

func InitTable(db *gorm.DB)  {
	DB = db
	db.AutoMigrate(&Account{})
}


// 当前数据库有多少条


