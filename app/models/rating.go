package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type Rating struct {
	ID        uint64         `gorm:"primary_key" json:"id"`
	DepotID   uint64         `json:"depot_id"`
	UserID    uint           `json:"user_id"`
	Komentar  string         `json:"komentar" gorm:"type:varchar(255)"`
	Rating    float64        `json:"rating"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relasi
	User  User  `json:"user" gorm:"foreignkey:UserID"`
	Depot Depot `json:"depot" gorm:"foreignkey:DepotID"`
}

func (Rating) TableName() string {
	return config.GetDBPrefix("rating")
}
