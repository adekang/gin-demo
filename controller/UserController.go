package controller

import (
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10)
	}

	log.Println(name, password, telephone)
	// 判断手机号是否存在

	if isTelExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户存在不允许注册",
		})
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0

}
