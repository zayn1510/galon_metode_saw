package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpRating(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.Rating{})
}

func DownRating(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.Rating{})
}
