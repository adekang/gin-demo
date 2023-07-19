package model

type MaterialAttr struct {
	AttrId int    `gorm:"primaryKey;autoIncrement;not null;column:attr_id"`
	Name   string `gorm:"type:varchar(255)"`
	Type   int    `gorm:"size:2"`
}

func (v MaterialAttr) TableName() string {
	return "material_attr"
}
