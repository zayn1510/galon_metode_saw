package resources

import (
	"github.com/zayn1510/goarchi/app/models"
)

type UsersResource struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email,omitempty"`
	Username           string `json:"username"`
	Password           string `json:"password,omitempty"`
	Role               string `json:"role"`
	Status             string `json:"status"`
	LastPasswordChange string `json:"last_password_change,omitempty"`
	NomorHandphone     string `json:"nomor_handphone"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func NewUserResource(m models.User) *UsersResource {
	var deletedAt string
	var lastPasswordChange string

	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	if m.LastPasswordChange != nil {
		lastPasswordChange = m.LastPasswordChange.Format("2006-01-02 15:04:05")
	}

	return &UsersResource{ // Mengembalikan pointer agar lebih ringan
		ID:                 m.ID,
		Name:               m.Nama,
		Email:              m.Email,
		Username:           m.Username,
		Password:           m.Password,
		Role:               m.Role,
		Status:             m.Status,
		LastPasswordChange: lastPasswordChange,
		NomorHandphone:     m.NomorHandphone,
		CreatedAt:          m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:          m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:          deletedAt,
	}
}

func GetUsersResource(data []models.User) []*UsersResource {
	resources := make([]*UsersResource, len(data))
	for i, v := range data {
		resources[i] = NewUserResource(v)
	}
	return resources
}
