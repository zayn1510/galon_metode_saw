package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type KriteriaService struct {
	db *gorm.DB
}

func NewKriteriaService() *KriteriaService {
	return &KriteriaService{
		db: config.GetDB(),
	}
}

func (s *KriteriaService) FindAll(offset, limit int, filter string) ([]models.Kriteria, error) {
	var resutl []models.Kriteria
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	query := s.db.Offset(offset).Limit(limit).Order("id asc")
	if filter != "" {
		query = query.Where("keterangan LIKE ?", "%"+filter+"%")
	}
	if err := query.Find(&resutl).Error; err != nil {
		return nil, err
	}
	return resutl, nil
}
func (s *KriteriaService) IsExistId(id uint) (*models.Kriteria, error) {
	var result models.Kriteria
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *KriteriaService) FindById(id uint) (*models.Kriteria, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *KriteriaService) Create(k *models.Kriteria) error {
	if err := s.db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

func (s *KriteriaService) Update(updateData *requests.UpdateKriteriaRequest, id uint) error {
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

func (s *KriteriaService) Delete(id uint) error {
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
