package controller

import (
	"github.com/gin-gonic/gin"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
)

// FindAllByUse 根据用途查找公式
func FindAllByUse(c *gin.Context) {

	useExp := c.Query("use")
	var expression []model.Expression

	if useExp == "" {
		response.Fail(c, nil, "查询字符0为空")
		return
	}

	db := common.GetDB()
	res := db.Where("expression.use = ?", useExp).Find(&expression)
	if res.Error != nil {
		response.Fail(c, nil, "查询失败")
		return
	}
	response.Success(c, expression, "查询成功")
}
