package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type UsersService struct {
	db *gorm.DB
}

func NewUsersService() *UsersService {
	return &UsersService{
		db: config.GetDB(),
	}
}

func (s *UsersService) FindAll(offset, limit int, filter string) ([]models.User, error) {
	var result []models.User
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	query := s.db.Offset(offset).Limit(limit).Order("id asc")
	if len(filter) > 0 {
		query = query.Where("name LIKE ?", "%"+filter+"%")
	}
	if err := query.Preload("UserLocation").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (s *UsersService) IsExistId(id uint) (*models.User, error) {
	var result models.User
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *UsersService) FindById(id uint) (*models.User, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UsersService) Create(k *models.User) error {
	if err := s.db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

func (s *UsersService) Update(updateData *requests.UpdateUserRequest, id uint) error {
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

func (s *UsersService) Delete(id uint) error {
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
