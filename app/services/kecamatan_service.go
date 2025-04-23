package services

import (
	"errors"
	"fmt"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
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

func (s *KecamatanService) FindAll(offset, limit int, filter string) ([]models.Kecamatan, error) {
	var resutl []models.Kecamatan
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	query := s.db.Offset(offset).Limit(limit).Order("id asc")
	if filter != "" {
		query = query.Where("nama_kecamatan LIKE ?", "%"+filter+"%")
	}
	if err := query.Find(&resutl).Error; err != nil {
		return nil, err
	}
	return resutl, nil
}

func (s *KecamatanService) JumlahDepotKecamatan() ([]*resources.JumlahDepotKecamatan, error) {
	var hasil []*resources.JumlahDepotKecamatan
	err := s.db.Raw(`
	SELECT 
		k.nama_kecamatan,
		COUNT(d.id) AS jumlah_depot
	FROM 
		tbl_kecamatan k
	LEFT JOIN 
		tbl_depot d ON d.kecamatan_id = k.id
	WHERE 
		k.deleted_at IS NULL and d.deleted_at IS NULL
	GROUP BY 
		k.id, k.nama_kecamatan
	ORDER BY 
		jumlah_depot DESC
`).Scan(&hasil).Error

	if err != nil {
		return nil, err
	}
	return hasil, nil
}

func (s *KecamatanService) CountAll() (int64, error) {
	var count int64
	err := s.db.Model(&models.Kecamatan{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
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
