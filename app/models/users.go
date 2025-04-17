package models

import (
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                 uint            `gorm:"primaryKey" json:"id"`
	Nama               string          `gorm:"type:varchar(100);not null" json:"nama"`
	Username           string          `gorm:"type:varchar(100);not null;" json:"username"`
	Email              string          `gorm:"type:varchar(100);not null" json:"email"`
	Password           string          `gorm:"type:varchar(255);not null" json:"-"`
	Role               string          `gorm:"type:enum('admin','user');default:'user'" json:"role"`
	Status             string          `gorm:"type:enum('active','inactive','banned');default:'active'" json:"status"`
	LastPasswordChange *time.Time      `json:"last_password_change"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
	DeletedAt          gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
	UserLocation       []User_location `gorm:"foreignkey:UserId" json:"user_location"`
}

func (User) TableName() string {
	return config.GetDBPrefix("users")
}
