package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type Depot struct {
	ID             uint           `gorm:"primary_key" json:"id"`
	KecamatanID    uint64         `json:"id_kecamatan"`
	NamaDepot      string         `json:"nama_depot" gorm:"type:varchar(100)"`
	Alamat         string         `json:"alamat" gorm:"type:varchar(100)"`
	Latitude       float64        `json:"latitude" gorm:"type:decimal(10,6)"`
	Longitude      float64        `json:"longitude" gorm:"type:decimal(10,6)"`
	NomorHandphone string         `json:"nomor_handphone" gorm:"type:varchar(20)"`
	Harga          int            `json:"harga"`
	Diskon         int            `json:"diskon"`
	Rating         float64        `json:"rating"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relasi
	Kecamatan Kecamatan `gorm:"foreignkey:KecamatanID" json:"kecamatan"`
}

func (Depot) TableName() string {
	return config.GetDBPrefix("depot")
}
