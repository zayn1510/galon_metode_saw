package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpUserLocation(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.User_location{})
}

func DownUserLocation(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.User_location{})
}
