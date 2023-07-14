package common

import (
	"fmt"
	"github/adekang/gin-demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gin_demo"
	username := "root"
	password := "123456"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	db.AutoMigrate(&model.User{})

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
