package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"healthy/config"
)

var Db *gorm.DB

func InitMySQL() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", config.AppConfig.AwsDbUsername, config.AppConfig.AwsDbPassword, config.AppConfig.AwsDbHost, "3306", "healthy")
	Db, err = gorm.Open(mysql.New(mysql.Config{
		//DSN:        "admin:1234@tcp(mysql80:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DSN: dsn,
	}), &gorm.Config{})
	return
}
