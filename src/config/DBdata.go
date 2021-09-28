package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
	"log"
)

var DBHelper *gorm.DB
var err error

func init()  {
	DBHelper, err = gorm.Open("mysql", "wxy:123456@tcp(127.0.0.1:3306)/gin-shop?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	log.Println("数据库运行...")
	DBHelper.LogMode(true)
	// 日志
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)
}