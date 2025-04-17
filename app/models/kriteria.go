package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type Kriteria struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Keterangan string `json:"keterangan" gorm:"size:255"`
	Bobot      float64
	Tipe       int            `json:"tipe" gorm:"default:0"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Kriteria) TableName() string {
	return config.GetDBPrefix("kriteria")
}
