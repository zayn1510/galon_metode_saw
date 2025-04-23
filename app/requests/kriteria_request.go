package requests

import "github.com/zayn1510/goarchi/app/models"

type CreateKriteriaRequest struct {
	Keterangan string  `json:"keterangan" validate:"required"`
	Bobot      float64 `json:"bobot" validate:"required"`
	Tipe       int     `json:"tipe,omitempty" `
}

type UpdateKriteriaRequest struct {
	Keterangan string  `json:"keterangan,omitempty" `
	Bobot      float64 `json:"bobot,omitempty"`
	Tipe       int     `json:"tipe;omitempty" `
}

func (r *CreateKriteriaRequest) ToModelKritera() *models.Kriteria {
	return &models.Kriteria{
		Keterangan: r.Keterangan,
		Bobot:      r.Bobot,
		Tipe:       r.Tipe,
	}
}
