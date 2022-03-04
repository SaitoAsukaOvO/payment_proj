package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"payment_proj/config"
	"payment_proj/model"
	"time"
)


var DB *gorm.DB


func init() {
	conf := config.LoadAppConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database)
	tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = tmp.AutoMigrate(&model.User{})
	_ = tmp.AutoMigrate(&model.Wallet{})
	_ = tmp.AutoMigrate(&model.Transaction{})
	sqlDB, _ := tmp.DB()
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetMaxIdleConns(15)
	DB = tmp
}