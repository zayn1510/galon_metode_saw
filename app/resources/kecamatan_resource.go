package resources

import "github.com/zayn1510/goarchi/app/models"

type KecamatanResource struct {
	ID            uint64 `json:"id"`
	NamaKecamatan string `json:"nama_kecamatan"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}
type JumlahDepotKecamatan struct {
	NamaKecamatan string `json:"nama_kecamatan"`
	JumlahDepot   int64  `json:"jumlah_depot"`
}

func NewKecamatanResource(m *models.Kecamatan) *KecamatanResource {
	var deleteAt string
	if m.DeletedAt.Valid {
		deleteAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &KecamatanResource{
		ID:            m.ID,
		NamaKecamatan: m.NamaKecamatan,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:     deleteAt,
	}
}

func GetKecamatanResource(data []models.Kecamatan) []*KecamatanResource {
	response := make([]*KecamatanResource, len(data))
	for i, _ := range data {
		response[i] = NewKecamatanResource(&data[i])
	}
	return response
}
