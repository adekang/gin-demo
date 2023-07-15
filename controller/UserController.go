package controller

import (
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/dto"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
	"github/adekang/gin-demo/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

// 登录
func Login(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码不正确,必须为11位")

		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {

		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)

	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")

		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

func Register(c *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	// 数据验证
	c.JSON(http.StatusOK, gin.H{
		"username": name,
		"password": password,
	})

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码不正确,必须为11位")

		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")

		return
	}
	// 名称不为空，为空返回10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, password, telephone)
	// 判断手机号是否存在
	if isTelExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户存在,s不允许注册")
	}

	// 密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	DB.Create(&newUser)

	// 返回结果
	response.Success(c, nil, "注册成功")
}

func isTelExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0

}
