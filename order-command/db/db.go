package db

import (
	"fmt"
	"go-cqrs-saga-edd/order-command/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("🚀 Failed to the Database, err message: ", err.Error())
		return nil
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
