package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type LoginLog struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	UserId    uint           `json:"user_id"`
	ISP       string         `json:"isp"`
	Username  string         `json:"username"`
	Nama      string         `json:"nama"`
	Role      string         `json:"role"`
	Status    string         `json:"status"`
	IPAddress string         `json:"ip_address"`
	Device    string         `json:"device"`
	Browser   string         `json:"browser"`
	Platform  string         `json:"platform"`
	Country   string         `json:"country"`
	City      string         `json:"city"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relasi
	User User `gorm:"foreignKey:UserId" json:"user"`
}

func (LoginLog) TableName() string {
	return config.GetDBPrefix("login_logs")
}
