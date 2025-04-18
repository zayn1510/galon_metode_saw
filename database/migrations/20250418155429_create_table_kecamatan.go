package migrations

import (
	"github.com/zayn1510/goarchi/app/models"
	"gorm.io/gorm"
)

func UpKecamatan(db *gorm.DB) error {
	// TODO: implement migration
	return db.AutoMigrate(&models.Kecamatan{})
}

func DownKecamatan(db *gorm.DB) error {
	// TODO: implement rollback
	return db.Migrator().DropTable(&models.Kecamatan{})

}
