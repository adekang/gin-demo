package model

import "gorm.io/gorm"

type TargetStatic struct {
	gorm.Model
	ExpId     int    `gorm:"size:24"`
	Connector string `gorm:"type:varchar(20)"`
	Used      string `gorm:"type:varchar(20)"`
	Value     string `gorm:"type:varchar(20)"`
}

func (v TargetStatic) TableName() string {
	return "target_static"
}
