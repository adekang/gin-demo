package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               args,
		DefaultStringSize: 171,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束 ,逻辑外键
	})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
