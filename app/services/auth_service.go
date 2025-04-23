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
