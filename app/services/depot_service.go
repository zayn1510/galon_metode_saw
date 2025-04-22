package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type DepotService struct {
	db *gorm.DB
}

func NewDepotService() *DepotService {
	return &DepotService{
		db: config.GetDB(),
	}
}
func (s *DepotService) CountAll() (int64, error) {
	var count int64
	err := s.db.Model(&models.Depot{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (s *DepotService) FindAll(offset, limit int, filter string) ([]models.Depot, error) {
	var resutl []models.Depot
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	query := s.db.Offset(offset).Limit(limit).Order("id asc")
	if filter != "" {
		query = query.Where("nama_depot LIKE ?", "%"+filter+"%")
	}
	if err := query.Find(&resutl).Error; err != nil {
		return nil, err
	}
	return resutl, nil
}

func (s *DepotService) IsExistId(id uint) (*models.Depot, error) {
	var result models.Depot
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *DepotService) FindById(id uint) (*models.Depot, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *DepotService) Create(k *models.Depot) error {
	if err := s.db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

func (s *DepotService) Update(data map[string]interface{}, id uint) error {
	// check id exist
	result, err := s.IsExistId(id)
	if err != nil {
		return err
	}

	if err := s.db.Model(result).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *DepotService) Delete(id uint) (*models.Depot, error) {
	// check id exist
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	if err := s.db.Delete(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
