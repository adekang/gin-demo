package controller

import (
	"github.com/gin-gonic/gin"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/dto"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
)

func FindAllAlphaBeta(c *gin.Context) {
	//连接数据库
	db := common.GetDB()

	err := db.AutoMigrate(&model.AlphaBeta{})
	if err != nil {
		return
	}

	//新增数据 到 AlphaBeta 表
	//u := model.AlphaBeta{
	//	Alpha: 1,
	//	Beta:  1,
	//	Apply: "go测试导热率",
	//}

	//if err := db.Create(&u).Error; err != nil {
	//	fmt.Println(err)
	//}

	///查询表
	var alphaBeta []model.AlphaBeta
	result := db.Find(&alphaBeta)
	// 为空 则返回数据
	if result.Error == nil {
		response.Success(c, gin.H{"result": dto.ToAlphaBetaDto(alphaBeta)}, "查询成功")
	} else {
		response.Fail(c, gin.H{}, "查询失败")
	}

}

func UpdateAlphaBeta(c *gin.Context) {
	db := common.GetDB()

	var requestAlphaBeta = model.AlphaBeta{}
	err := c.Bind(&requestAlphaBeta)
	if err != nil {
		response.Fail(c, gin.H{}, "请求错误")
		return
	}
	id := requestAlphaBeta.ID

	updateVar := &model.AlphaBeta{
		ID:    id,
		Alpha: requestAlphaBeta.Alpha,
		Beta:  requestAlphaBeta.Beta,
		Apply: requestAlphaBeta.Apply,
	}

	if err := db.Model(&requestAlphaBeta).Updates(&updateVar).Error; err != nil {
		// 处理保存失败错误
		response.Fail(c, gin.H{}, "更新失败")
		return

	}

	response.Success(c, gin.H{}, "更新成功")
}
