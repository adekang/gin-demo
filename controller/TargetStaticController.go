package controller

import (
	"github.com/gin-gonic/gin"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/dto"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
)

func FindAllTargetStatic(c *gin.Context) {
	db := common.GetDB()
	err := db.AutoMigrate(&model.Expression{}, &model.TargetStatic{})
	if err != nil {
		return
	}

	var targetStatic []model.TargetStatic
	result := db.Preload("Expression").Find(&targetStatic)

	if result.Error == nil {
		response.Success(c, dto.ToTargetStaticDto(targetStatic), "查询成功")
	} else {
		response.Fail(c, nil, "查询失败")
		return
	}
}

func UpdateTarget(c *gin.Context) {
	db := common.GetDB()

	var requestTargetStatic []model.TargetStatic
	err := c.Bind(&requestTargetStatic)
	if err != nil {
		response.Fail(c, nil, "请求错误")
		return
	}
	updateVar := &model.TargetStatic{}
	// 发来2条数据
	if len(requestTargetStatic) > 1 {
		updateVar.ID = requestTargetStatic[0].ID
		updateVar.ExpId = requestTargetStatic[0].ExpId
		updateVar.Used = requestTargetStatic[0].Used

		for i, item := range requestTargetStatic {
			exp := ""
			if i != len(requestTargetStatic)-1 {
				exp = ","
			}
			updateVar.Connector += item.Connector + exp
			updateVar.Value += item.Value + exp
		}
	} else {
		updateVar = &requestTargetStatic[0]
	}

	if err := db.Model(&requestTargetStatic).Updates(&updateVar).Error; err != nil {
		// 处理保存失败错误
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}
