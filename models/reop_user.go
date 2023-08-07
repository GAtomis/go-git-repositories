package models

type RepoUser struct {
	Rid  int `gorm:"column:rid;type:int(11);" json:"rid"`      //仓库id
	Uid  int `gorm:"column:uid;type:int(11);" json:"uid"`      //仓库id
	Type int `gorm:"column:type;type:tinyint(1);" json:"type"` // 类型 1 所有者 2被授权者」

}

func (ru *RepoUser) TableName() string {
	return "repo_user"
}
