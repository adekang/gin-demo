package model

import "gorm.io/gorm"

type AlphaBeta struct {
	gorm.Model
	ID    int     `gorm:"primary_key;AUTO_INCREMENT"`
	Alpha float32 `gorm:"size:24"`
	Beta  float32 `gorm:"size:24"`
	Apply string  `gorm:"type:varchar(20);not null"`
}
