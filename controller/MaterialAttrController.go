package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/adekang/gin-demo/common"
	"github/adekang/gin-demo/model"
	"github/adekang/gin-demo/response"
	"strconv"
	"strings"
)

func FindAllMaterialAttr(c *gin.Context) {
	db := common.GetDB()
	err := db.AutoMigrate(&model.MaterialAttr{})
	if err != nil {
		return
	}

	var materialAttr []model.MaterialAttr
	result := db.Find(&materialAttr)
	// 为空 则返回数据
	if result.Error == nil {
		response.Success(c, materialAttr, "查询成功")
	} else {
		response.Fail(c, gin.H{}, "查询失败")
	}

}

func UpdateMaterialAttr(c *gin.Context) {

	requestMaterialAttr := model.MaterialAttr{}
	err := c.Bind(&requestMaterialAttr)
	if err != nil {
		response.Fail(c, gin.H{}, "请求错误")
		return
	}
	db := common.GetDB()
	var materialAttr model.MaterialAttr

	if err := db.Model(&materialAttr).Where("attr_id = ?", requestMaterialAttr.AttrId).Updates(&requestMaterialAttr).Error; err != nil {
		// 处理保存失败错误
		response.Fail(c, gin.H{}, "更新失败")
		return
	}
	response.Success(c, gin.H{}, "更新成功")

}

func DeleteMaterialAttr(c *gin.Context) {
	// 连接数据库
	db := common.GetDB()
	//	获取请求体
	requestMaterialAttr := model.MaterialAttr{}
	err := c.Bind(&requestMaterialAttr)
	if err != nil {
		response.Fail(c, gin.H{}, "请求错误")
		return
	}

	//	根据id删除
	var materialAttr model.MaterialAttr

	if err := db.Where("attr_id = ?", requestMaterialAttr.AttrId).Delete(&materialAttr).Error; err != nil {
		// 处理保存失败错误
		response.Fail(c, gin.H{}, "删除成功")
		return
	}
	response.Success(c, gin.H{}, "更新成功")
}

type Request struct {
	Eg float32 `json:"eg"`
}

// ScreenMaterialAttr 判断eg  获取 粘结剂和单质炸药
func ScreenMaterialAttr(c *gin.Context) {
	// Query 获取参数
	requestQuery := c.Query("eg")
	f, err := strconv.ParseFloat(requestQuery, 32)
	if err != nil {
		fmt.Println("无法将字符串转换为float32:", err)
		return
	}
	eg := float32(f)
	//	expression表中的 name和use筛选出整条数据
	var expression model.Expression
	db := common.GetDB()
	db.Model(&model.Expression{}).Where("name = ? AND expression.use = ?", "Eg", "筛选").Find(&expression)

	//	获取筛选出来的 exp_id
	expId := expression.ExpId
	//	根据 exp_id 在 targetStatic 表中 获取 connector 和 value
	var targetStatic model.TargetStatic
	db.Unscoped().Where("exp_id = ?", expId).Find(&targetStatic)

	connector := targetStatic.Connector
	value := targetStatic.Value
	// judge 判断 eg的正确性
	if judges(connector, value, eg) {
		//	正确返回 粘结剂和单质炸药

		//连接剂
		var connectionAgent []model.MaterialInfo
		db.Where("m_category = ?", "连接剂").Find(&connectionAgent)

		//单质炸药
		dynamite := model.MaterialInfo{}
		db.Where("m_category = ?", "单质").Find(&dynamite)

		//	返回数据
		response.Success(c, gin.H{
			"simSub": dynamite,
			"linker": connectionAgent,
		}, "查询成功")

	} else {
		//	不正确返回字符串
		alphaBeta := model.AlphaBeta{}
		// GORM自动生成的SQL语句包含了deleted_at条件,导致未知列错误。
		// 而数据表中实际上并没有 deleted_at
		// 通过Unscoped()方法可以解决这个问题 或者在类型中直接去掉gorm.Model
		db.Unscoped().Where("apply = ?", "圆筒比动能").First(&alphaBeta)

		alpha := alphaBeta.Alpha
		beta := alphaBeta.Beta
		text := "现材料不满足设计要求,需研发Eg>" +
			formatFloat(eg/(0.92*alpha)) +
			"且H50>" +
			formatFloat(0.32*beta) +
			"cm的新型单质炸药"
		response.Fail(c, gin.H{}, text)
	}

}

// 格式化float32 为 string
func formatFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

// 判断 eg 是否满足条件
func judges(connectors string, values string, eg float32) bool {

	cons := strings.Split(connectors, ",")
	vals := strings.Split(values, ",")

	for i := 0; i < len(cons); i++ {
		if !judge(cons[i], vals[i], eg) {
			return false
		}
	}
	return true
}

// 真正判断 eg 是否满足条件
func judge(connector string, value string, eg float32) bool {
	valueNum, err := strconv.ParseFloat(value, 32)
	if err != nil {
		// 处理错误
	}
	if connector == ">" {
		if eg > float32(valueNum) {
			return true
		} else {
			return false
		}
	} else if connector == "<" {
		if eg < float32(valueNum) {
			return true
		} else {
			return false
		}
	} else if connector == ">=" {
		if eg >= float32(valueNum) {
			return true
		} else {
			return false
		}
	} else if connector == "<=" {
		if eg <= float32(valueNum) {
			return true
		} else {
			return false
		}
	} else {
		if eg == float32(valueNum) {
			return true
		} else {
			return false
		}
	}
}
