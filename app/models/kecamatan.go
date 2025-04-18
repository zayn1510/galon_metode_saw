package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type Kecamatan struct {
	ID            uint64         `gorm:"primary_key" json:"id"`
	NamaKecamatan string         `json:"nama_kecamatan"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Kecamatan) TableName() string {
	return config.GetDBPrefix("kecamatan")
}
