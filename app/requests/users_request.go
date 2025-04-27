package requests

import (
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/core/tools"
	"time"
)

type CreateUserRequest struct {
	Nama           string `json:"nama" validate:"required"`
	Email          string `json:"email,omitempty"`
	Username       string `json:"username" validate:"required,min=3,max=100,alphanum"`
	Password       string `json:"password" validate:"required,min=8"`
	Role           string `json:"role" validate:"omitempty,oneof=admin user"`
	NomorHandphone string `json:"nomor_handphone" validate:"required"`
	Status         string `json:"status" validate:"omitempty,oneof=active inactive banned"`
}

type UpdateUserRequest struct {
	Nama               *string    `json:"nama,omitempty"`
	Email              *string    `json:"email,omitempty"`
	Username           *string    `json:"username,omitempty"`
	Password           *string    `json:"password,omitempty"`
	Role               *string    `json:"role,omitempty" validate:"omitempty,oneof=admin user"`
	Status             *string    `json:"status,omitempty" validate:"omitempty,oneof=active inactive banned"`
	NomorHandphone     *string    `json:"nomor_handphone,omitempty"`
	LastPasswordChange *time.Time `json:"last_password_change,omitempty"`
}

func (req *CreateUserRequest) ToUser() (*models.User, error) {
	hasher := &tools.BcryptHasher{}
	hash, err := hasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username:       req.Username,
		Nama:           req.Nama,
		Email:          req.Email,
		Password:       hash,
		Role:           req.Role,
		Status:         req.Status,
		NomorHandphone: req.NomorHandphone,
	}, nil
}
