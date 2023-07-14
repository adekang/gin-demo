package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

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

	db.AutoMigrate(&User{})

	return db

}

func isTelExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0

}

func main() {

	db := InitDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/auth/register", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")
		c.JSON(http.StatusOK, gin.H{
			"username": name,
			"password": password,
		})
		log.Println(telephone, len(telephone))

		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号码不正确,必须为11位",
			})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "密码不能少于6位",
			})
			return
		}
		// 名称不为空，为空返回10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		log.Println(name, password, telephone)
		// 判断手机号是否存在

		if isTelExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "用户存在不允许注册",
			})
		}

		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&newUser)

		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)

}
