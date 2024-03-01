// Package database 连接mysql数据库
package database

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 读取.ini里面的数据库配置
	config, err1 := ini.Load("./config/app.ini")
	if err1 != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	// dsn := "root:123456@tcp(192.168.0.6:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, // 打印sql
		// SkipDefaultTransaction: true, //禁用事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
