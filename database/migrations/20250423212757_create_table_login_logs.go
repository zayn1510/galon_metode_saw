package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpLoginLogs(db *gorm.DB) error {
	return db.AutoMigrate(&models.LoginLog{})
}

func DownLoginLogs(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.LoginLog{})

}
