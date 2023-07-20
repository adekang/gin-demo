package model

// Expression 公式
type Expression struct {
	ExpId      int    `gorm:"primaryKey;AUTO_INCREMENT;column:exp_id" json:"expId"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Use        string `gorm:"type:varchar(255)" json:"use"`
	Expression string `gorm:"type:varchar(255)" json:"expression"`
}

func (v Expression) TableName() string {
	return "expression"
}
