package requests

import "github.com/zayn1510/goarchi/app/models"

type CreateRatingRequest struct {
	UserID   uint    `json:"user_id" validate:"required,numeric,min=1"`
	DepotID  uint64  `json:"depot_id" validate:"required,numeric,min=1"`
	Komentar string  `json:"komentar" validate:"required"`
	Rating   float64 `json:"rating" validate:"required"`
}

type UpdateRatingRequest struct {
	UserID   *uint    `json:"user_id,omitempty"`
	DepotID  uint64   `json:"depot_id,omitempty"`
	Komentar *string  `json:"komentar,omitempty"`
	Rating   *float64 `json:"rating,omitempty"`
}

func (req *CreateRatingRequest) ToModelRating() *models.Rating {
	return &models.Rating{
		UserID:   req.UserID,
		DepotID:  req.DepotID,
		Komentar: req.Komentar,
		Rating:   req.Rating,
	}
}
