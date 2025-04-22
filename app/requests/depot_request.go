package requests

import (
	"github.com/zayn1510/goarchi/app/models"
	"mime/multipart"
)

type CreateDepotRequest struct {
	NamaDepot      string                `form:"nama_depot" validate:"required"`
	Alamat         string                `form:"alamat" validate:"required"`
	Latitude       float64               `form:"latitude" validate:"required"`
	Longitude      float64               `form:"longitude" validate:"required"`
	NomorHandphone string                `form:"nomor_handphone" validate:"required"`
	Harga          int                   `form:"harga" validate:"required"`
	Diskon         int                   `form:"diskon" validate:"required"`
	Rating         float64               `form:"rating" validate:"required"`
	KecamatanId    uint64                `form:"kecamatan_id" validate:"required"`
	Foto           *multipart.FileHeader `form:"foto"`
}

type UpdateDepotRequest struct {
	NamaDepot      *string               `form:"nama_depot"`
	Alamat         *string               `form:"alamat"`
	Latitude       *float64              `form:"latitude"`
	Longitude      *float64              `form:"longitude"`
	NomorHandphone *string               `form:"nomor_handphone"`
	Harga          *int                  `form:"harga"`
	Diskon         *int                  `form:"diskon"`
	Rating         *float64              `form:"rating"`
	KecamatanId    *uint64               `form:"kecamatan_id"`
	Foto           *multipart.FileHeader `form:"foto"`
	FotoLama       string                `form:"foto_lama"`
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
