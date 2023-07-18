package model

import "gorm.io/gorm"

type TargetStatic struct {
	gorm.Model
	Connector  string     `gorm:"type:varchar(20)"`
	Used       string     `gorm:"type:varchar(20)"`
	Value      string     `gorm:"type:varchar(20)"`
	Expression Expression `gorm:"embedded"` // 嵌套字段
	ExpId      int        `gorm:"foreign_Key:ExpressionExpId"`
}

func (v TargetStatic) TableName() string {
	return "target_static"
}
