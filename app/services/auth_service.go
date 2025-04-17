package services

import "github.com/zayn1510/goarchi/core/tools"

type AuthService struct {
	PasswordHasher *tools.PasswordHasher
}

func newAuthService(passwordHasher *tools.PasswordHasher) *AuthService {
	return &AuthService{
		PasswordHasher: passwordHasher,
	}
}
