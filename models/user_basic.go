package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);"json:"identity"`
	Username string `gorm:"column:username;type:varchar(255); uniqueIndex;"json:"username"`
	Password string `gorm:"column:password;type:varchar(36);"json:"password"`
}

func (u *UserBasic) TableName() string {
	return "user_basic"

}
