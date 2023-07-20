package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
)

func FindStaticByType(c *gin.Context) {
	requestQuery := c.Query("type")

	var targetStatic, resultTarget []model.TargetStatic

	db := common.GetDB()
	db.Unscoped().Where("used = ?", requestQuery).Find(&targetStatic)

	// 根据exp_id 获得 相关信息，然后进行拼接
	for i, item := range targetStatic {
		expId := item.ExpId
		var expression model.Expression
		db.Where("exp_id = ?", expId).Find(&expression)
		targetStatic[i].Expression = expression
		resultTarget = append(resultTarget, targetStatic[i])

	}
	fmt.Println(resultTarget)
	response.Success(c, resultTarget, "查询成功")
}
