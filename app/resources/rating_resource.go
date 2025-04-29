package resources

import (
	"github.com/zayn1510/goarchi/app/models"
	"math"
)

type RatingResource struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	DepotID   uint64 `json:"depot_id"`
	Nama      string `json:"nama"`
	Depot     string `json:"depot"`
	Komentar  string `json:"komentar,omitempty"`
	Rating    int    `json:"rating"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func NewRatingResource(m models.Rating) *RatingResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &RatingResource{ // Mengembalikan pointer agar lebih ringan
		ID:        m.ID,
		UserID:    uint64(m.UserID),
		DepotID:   m.DepotID,
		Nama:      m.User.Nama,
		Depot:     m.Depot.NamaDepot,
		Komentar:  m.Komentar,
		Rating:    int(math.Floor(m.Rating)),
		Role:      m.User.Role,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func GetRatingResource(data []*models.Rating) []*RatingResource {
	resources := make([]*RatingResource, len(data))
	for i, v := range data {
		resources[i] = NewRatingResource(*v)
	}
	return resources
}
