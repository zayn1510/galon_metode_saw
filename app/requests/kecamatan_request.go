package requests

import "github.com/zayn1510/goarchi/app/models"

type CreatedKecamatanRequest struct {
	NamaKecamatan string `json:"nama_kecamatan" validate:"required"`
}
type UpdatedKecamatanRequest struct {
	NamaKecamatan string `json:"nama_kecamatan,omitempty"`
}

func (req *CreatedKecamatanRequest) ToKecamatan() *models.Kecamatan {
	return &models.Kecamatan{
		NamaKecamatan: req.NamaKecamatan,
	}
}
