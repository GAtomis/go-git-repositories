package models

import "gorm.io/gorm"

type RepoBasic struct {
	gorm.Model
	Uid      int    `gorm:"column:uid;type:int(11);" json:"uid"`
	Identity string `gorm:"column:identity;type:varchar(36);"json:"identity"`
	Name     string `gorm:"column:name;type:varchar(255);" json:"name" `
	Desc     string `gorm:"column:name;type:varchar(255);" json:"desc" `
	Star     int    `gorm:"column:desc;type:int(11);" json:"star" `
}

func (table *RepoBasic) TableName() string {
	return "repo_basic"
}
