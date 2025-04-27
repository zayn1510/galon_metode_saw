package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type UserlocationService struct {
	db *gorm.DB
}

func NewUserLocationService() *UserlocationService {
	return &UserlocationService{
		db: config.GetDB(),
	}
}

func (s *UserlocationService) FindAll(offset, limit int) ([]models.User_location, error) {
	var resutl []models.User_location
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	if err := s.db.Offset(offset).Limit(limit).Order("id asc").Preload("User").Find(&resutl).Error; err != nil {
		return nil, err
	}
	return resutl, nil
}

func (s *UserlocationService) IsExistId(id uint) (*models.User_location, error) {
	var result models.User_location
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}

func (s *UserlocationService) IsExistIdUser(user_id uint) (*models.User_location, error) {
	var result models.User_location
	if err := s.db.Where("user_id =?", user_id).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}

func (s *UserlocationService) IsUserExist(userid uint) (*models.User_location, error) {
	var result models.User_location
	if err := s.db.Order("id desc").Where("user_id = ?", userid).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("UserID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *UserlocationService) FindById(id uint) (*models.User_location, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserlocationService) Create(k *models.User_location) error {
	if err := s.db.Create(k).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserlocationService) Update(updateData *requests.UpdateUserLocationRequest, id uint) error {
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

func (s *UserlocationService) Delete(id uint) error {
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
