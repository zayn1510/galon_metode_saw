package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpKriteria(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.Kriteria{})
}

func DownKriteria(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.Kriteria{})
}
