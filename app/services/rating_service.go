package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"gorm.io/gorm"
)

type RatingService struct {
	db *gorm.DB
}

func NewRatingService() *RatingService {
	return &RatingService{
		db: config.GetDB(),
	}
}

func (s *RatingService) FindAll(offset, limit int, depotId uint) ([]*models.Rating, error) {
	var result []*models.Rating
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	query := s.db.Offset(offset).Limit(limit).Order("id desc").Preload("User").Preload("Depot")

	if depotId > 0 {
		query = query.Where("depot_id = ?", depotId)
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *RatingService) IsExistId(id uint) (*models.Rating, error) {
	var result models.Rating
	if err := s.db.First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("ID tidak ditemukan")
		}
		return nil, err
	}
	return &result, nil
}
func (s *RatingService) FindById(id uint) (*models.Rating, error) {
	result, err := s.IsExistId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *RatingService) Create(k *models.Rating) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	depot, err := NewDepotService().FindById(uint(k.DepotID))
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(k).Error; err != nil {
		tx.Rollback()
		return err
	}

	var result struct {
		Count int64
		Sum   float64
	}
	if err := tx.Model(&models.Rating{}).
		Select("COALESCE(SUM(rating),0) as sum, COUNT(*) as count").
		Where("depot_id = ?", k.DepotID).Scan(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("depot tidak ditemukan")
		}
		return err
	}

	depot.Rating = result.Sum / float64(result.Count)
	if err := tx.Save(&depot).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit jika semua berhasil
	return tx.Commit().Error
}

func (s *RatingService) Update(updateData *requests.UpdateRatingRequest, id uint) error {
	// check id exist
	result, err := s.IsExistId(id)
	if err != nil {
		return err
	}
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	depot, err := NewDepotService().FindById(uint(updateData.DepotID))
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(result).Updates(updateData).Error; err != nil {
		return err
	}

	var rating struct {
		Count int64
		Sum   float64
	}
	if err := tx.Model(&models.Rating{}).
		Select("COALESCE(SUM(rating),0) as sum, COUNT(*) as count").
		Where("depot_id = ?", updateData.DepotID).Scan(&rating).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("depot tidak ditemukan")
		}
		return err
	}

	depot.Rating = rating.Sum / float64(rating.Count)
	if err := tx.Save(&depot).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *RatingService) Delete(id uint) error {
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
