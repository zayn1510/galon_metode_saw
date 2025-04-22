package resources

import (
	"github.com/zayn1510/goarchi/app/models"
)

type KriteriaResource struct {
	ID        uint    `json:"id"`
	Kriteria  string  `json:"keterangan"`
	Bobot     float64 `json:"bobot"`
	Tipe      int     `json:"tipe"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt string  `json:"deletedAt"`
}

func NewKriteriaResource(m models.Kriteria) *KriteriaResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &KriteriaResource{ // Mengembalikan pointer agar lebih ringan
		ID:        m.ID,
		Kriteria:  m.Keterangan,
		Bobot:     m.Bobot,
		Tipe:      m.Tipe,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func GetKriteriaResource(data []models.Kriteria) []*KriteriaResource {
	resources := make([]*KriteriaResource, len(data))
	for i, v := range data {
		resources[i] = NewKriteriaResource(v)
	}
	return resources
}
