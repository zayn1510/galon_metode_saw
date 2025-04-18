package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type KecamatanService struct {
	db *gorm.DB
}

func NewKecamatanService() *KecamatanService {
	return &KecamatanService{
		db: config.GetDB(),
	}
}

func (s *KecamatanService) FindAll(offset, limit int) ([]models.Kecamatan, error) {
	var resutl []models.Kecamatan
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	if err := s.db.Offset(offset).Limit(limit).Order("id asc").Find(&resutl).Error; err != nil {
		return nil, err
	}
	return resutl, nil
}
func (s *KecamatanService) IsExistId(id uint) (*models.Kecamatan, error) {
	var result models.Kecamatan
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *KecamatanService) FindById(id uint) (*models.Kecamatan, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *KecamatanService) Create(k *models.Kecamatan) error {
	if err := s.db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

func (s *KecamatanService) Update(updateData *requests.UpdatedKecamatanRequest, id uint) error {
	// check id exist
	result, err := s.IsExistId(id)
	if err != nil {
		return err
	}
	if err := s.db.Model(result).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (s *KecamatanService) Delete(id uint) error {
	// check id exist
	result, err := s.IsExistId(id)
	if err != nil {
		return err
	}
	if err := s.db.Delete(result).Error; err != nil {
		return err
	}
	return nil
}
