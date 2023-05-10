package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	fmt.Println("初始化链接")
	createDbConnection()
}
func createDbConnection() {
	dbConfig, err := ParseDb("resources/env.ini")
	if err != nil {
		fmt.Println("解析配置文件异常")
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.UserName, dbConfig.Password, dbConfig.Url, dbConfig.Port, dbConfig.DbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
}
