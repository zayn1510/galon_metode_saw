package services

import (
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type StatsService struct {
	db *gorm.DB
}

func NewStatsService() *StatsService {
	return &StatsService{
		db: config.GetDB(),
	}
}

func (s *StatsService) GetStatistik() *resources.StatsResource {
	var userCount, kriteriaCount, kecamatanCount, depotCount int64

	s.db.Model(&models.User{}).Count(&userCount)
	s.db.Model(&models.Kriteria{}).Count(&kriteriaCount)
	s.db.Model(&models.Kecamatan{}).Count(&kecamatanCount)
	s.db.Model(&models.Depot{}).Count(&depotCount)

	return &resources.StatsResource{
		Kriteria:  kriteriaCount,
		Users:     userCount,
		Kecamatan: kecamatanCount,
		Depot:     depotCount,
	}
}
