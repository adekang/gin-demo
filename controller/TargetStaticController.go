package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
)

func FindAllTargetStatic(c *gin.Context) {
	db := common.GetDB()
	err := db.AutoMigrate(&model.Expression{}, &model.TargetStatic{})
	if err != nil {
		return
	}

	var (
		targetStatic []model.TargetStatic
	)

	result := db.Preload("Expression").Find(&targetStatic)

	fmt.Println("--------------------------")
	fmt.Println(targetStatic)
	fmt.Println("--------------------------")

	if result.Error == nil {
		//response.Success(c, dto.ToTargetStaticDto(targetStatic), "查询成功")
		response.Success(c, gin.H{
			"result": targetStatic,
		}, "查询成功")
	} else {
		response.Fail(c, gin.H{}, "查询失败")
		return
	}
}

func UpdateTarget(c *gin.Context) {
	db := common.GetDB()

	var requestTargetStatic = model.TargetStatic{}
	err := c.Bind(&requestTargetStatic)
	if err != nil {
		response.Fail(c, gin.H{}, "请求错误")
		return
	}
	updateVar := &model.TargetStatic{}
	err = copier.Copy(updateVar, requestTargetStatic)

	if err != nil {
		// 处理错误
	}
	if err := db.Model(&requestTargetStatic).Updates(&updateVar).Error; err != nil {
		// 处理保存失败错误
		response.Fail(c, gin.H{}, "更新失败")
		return
	}

	response.Success(c, gin.H{}, "更新成功")
}
