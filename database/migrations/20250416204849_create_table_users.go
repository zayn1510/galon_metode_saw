package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpUsers(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.User{})
}

func DownUsers(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.User{})
}
