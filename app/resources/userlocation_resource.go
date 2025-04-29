package resources

import "github.com/zayn1510/goarchi/app/models"

type UserlocationResource struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Username  string  `json:"username,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}

func NewUserLocationResource(m models.User_location) *UserlocationResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &UserlocationResource{ // Mengembalikan pointer agar lebih ringan
		ID:        m.ID,
		UserID:    uint(m.UserId),
		Name:      m.User.Nama,
		Username:  m.User.Username,
		Longitude: m.Longitude,
		Latitude:  m.Latitude,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func GetUserLocationResource(data []models.User_location) []*UserlocationResource {
	resources := make([]*UserlocationResource, len(data))
	for i, v := range data {
		resources[i] = NewUserLocationResource(v)
	}
	return resources
}
