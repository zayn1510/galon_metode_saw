package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type User_location struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	UserId    uint64         `json:"user_id"`
	Latitude  float64        `json:"latitude" gorm:"type:decimal(10,6)"`
	Longitude float64        `json:"longitude" gorm:"type:decimal(10,6)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	User User `gorm:"foreignKey:UserId" json:"user"`
}

func (User_location) TableName() string {
	return config.GetDBPrefix("user_location")
}
