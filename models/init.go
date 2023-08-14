package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlType struct {
	db *gorm.DB
}

const (
	SQL_URL = "43.143.244.182"
	// SQL_URL    = "127.0.0.1"
	SQL_NAME   = "go_git_repositories"
	SQL_CONFIG = "?charset=utf8mb4&parseTime=True&loc=Local"
	USERNAME   = "root"
	PASSWORD   = "QSCesz2rfx@"
	SQL_PORT   = "3306"
	PROTOCOL   = "tcp"
	TIME_OUT   = "30s"
)

var GAA_SQL = new(SqlType)

func NewDB() (db *gorm.DB) {
	return GAA_SQL.GetDB()
}

func (sql *SqlType) StartSQL() (db *gorm.DB, err error) {

	dsn := GetDsn()

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接异常,%s", err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&UserBasic{})
	sql.db = db

	return db, err
}
func (sql *SqlType) GetDB() (db *gorm.DB) {

	return sql.db

}

/**
 * @description: 获取sql连接
 * @return {*}
 * @Date: 2022-07-20 14:06:48
 */
func GetDsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s&%s", USERNAME, PASSWORD, PROTOCOL, SQL_URL, SQL_PORT, SQL_NAME, SQL_CONFIG, TIME_OUT)
	fmt.Println(dsn)
	return
}
