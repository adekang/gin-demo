package model

import "gorm.io/gorm"

type TargetStatic struct {
	gorm.Model
	Connector  string     `gorm:"type:varchar(20)"`
	Used       string     `gorm:"type:varchar(20)"`
	Value      string     `gorm:"type:varchar(20)"`
	ExpId      int        `gorm:"column:exp_id;foreignKey:ExpId"`
	Expression Expression `gorm:"foreignKey:ExpId;references:exp_id"`
}

func (v TargetStatic) TableName() string {
	return "target_static"
}
