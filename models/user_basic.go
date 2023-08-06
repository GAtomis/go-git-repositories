package models

type UserBasic struct {
	Identity string `gorm:"column:identity;type:varchar(36);"json:"identity"`
	Username string `gorm:"column:username;type:varchar(255);"json:"username"`
	Password string `gorm:"column:password;type:varchar(36);"json:"password"`
}

func (u *UserBasic) TableName() string {
	return "user_basic"

}
