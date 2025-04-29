package services

import (
	"errors"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/config"
	"github.com/zayn1510/goarchi/core/tools"
	"gorm.io/gorm"
	"log"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		db: config.GetDB(),
	}
}

func (s *AuthService) Login(req *requests.AuthRequest) (*models.User, error) {
	var user models.User
	hasher := &tools.BcryptHasher{}

	if err := s.db.Model(&models.User{}).Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Login failed for username '%s': user not found", req.Username)
			return nil, errors.New("invalid credentials")
		}
		log.Printf("Login failed for username '%s': user not found", req.Username)
		return nil, errors.New("invalid credentials") // hindari error spesifik untuk security
	}

	if err := hasher.Compare(user.Password, req.Password); err != nil {
		log.Printf("Login failed for username '%s': wrong password", req.Username)
		return nil, errors.New("invalid credentials")
	}

	// Sukses login
	log.Printf("Login success for username '%s'", req.Username)
	return &user, nil
}

func (s *AuthService) SaveLoginLogs(loginlog *models.LoginLog) error {
	if err := s.db.Save(loginlog).Error; err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GetLoginLogs(offset, limit int, filter string) ([]models.LoginLog, error) {
	var result []models.LoginLog
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
	if err := query.Preload("User").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (s *AuthService) CountLoginLogs() (int64, error) {
	var count int64
	err := s.db.Model(&models.LoginLog{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *AuthService) UpdatePassword(req *requests.UpdatePasswordRequest) error {
	hasher := &tools.BcryptHasher{}
	authReq := &requests.AuthRequest{
		Username: req.Username,
		Password: req.Password,
	}
	user, err := s.Login(authReq)
	if err != nil {
		return err
	}
	hashedPassword, err := hasher.Hash(req.NewPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := s.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}
