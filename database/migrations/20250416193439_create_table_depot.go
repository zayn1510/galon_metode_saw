package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpDepot(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.Depot{})

}

func DownDepot(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.Depot{})

}
