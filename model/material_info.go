package model

type MaterialInfo struct {
	MatId     int    `gorm:"primaryKey;autoIncrement;not null;column:mat_id"`
	Name      string `gorm:"type:varchar(255)"`
	MCategory string `gorm:"type:varchar(255)"`
}

func (v MaterialInfo) TableName() string {
	return "material_info"
}
