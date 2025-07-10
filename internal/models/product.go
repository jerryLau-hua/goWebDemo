package models

import "gorm.io/gorm"

//集成gorm.Model

type Product struct {
	gorm.Model         // 嵌入gorm.Model  id, created_at, updated_at, deleted_at
	Name       string  `gorm:"size:255"`
	Price      float64 `gorm:"type:decimal(10,2)"`
	Stock      int     `gorm:"default:0"`
}
