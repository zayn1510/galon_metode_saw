package requests

import "github.com/zayn1510/goarchi/app/models"

type CreateDepotRequest struct {
	NamaDepot      string  `json:"nama_depot" validate:"required"`
	Alamat         string  `json:"alamat" validate:"required"`
	Latitude       float64 `json:"latitude" validate:"required"`
	Longitude      float64 `json:"longitude" validate:"required"`
	NomorHandphone string  `json:"nomor_handphone" validate:"required"`
	Harga          int     `json:"harga" validate:"required"`
	Diskon         int     `json:"diskon" validate:"required"`
	Rating         float64 `json:"rating" validate:"required"`
	KecamatanId    uint64  `json:"kecamatan_id" validate:"required"`
}

type UpdateDepotRequest struct {
	NamaDepot      *string  `json:"nama_depot,omitempty"`
	Alamat         *string  `json:"alamat,omitempty"`
	Latitude       *float64 `json:"latitude,omitempty"`
	Longitude      *float64 `json:"longitude,omitempty"`
	NomorHandphone *string  `json:"nomor_handphone,omitempty"`
	Harga          *int     `json:"harga,omitempty"`
	Diskon         *int     `json:"diskon,omitempty"`
	Rating         *float64 `json:"rating,omitempty"`
	KecamatanId    *uint64  `json:"kecamatan_id,omitempty"`
}

func (req *CreateDepotRequest) ToDepot() *models.Depot {
	depot := &models.Depot{
		NamaDepot:      req.NamaDepot,
		Alamat:         req.Alamat,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		NomorHandphone: req.NomorHandphone,
		Harga:          req.Harga,
		Diskon:         req.Diskon,
		Rating:         req.Rating,
		KecamatanID:    req.KecamatanId,
	}
	return depot
}
