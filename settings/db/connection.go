package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB接続設定
func ConnectDB() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "wordbook"

	// datetime型をtime.Timeで受け取れるようにするため,parseTime=trueを指定している
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	return gorm.Open(DBMS, CONNECT)
}