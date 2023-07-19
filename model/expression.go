package model

// Expression 公式
type Expression struct {
	ExpId      int    `gorm:"primaryKey;AUTO_INCREMENT;column:exp_id"`
	Name       string `gorm:"type:varchar(255)"`
	Use        string `gorm:"type:varchar(255)"`
	Expression string `gorm:"type:varchar(255)"`
}

func (v Expression) TableName() string {
	return "expression"
}
