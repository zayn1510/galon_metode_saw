package requests

import "github.com/zayn1510/goarchi/app/models"

type CreateUserLocationRequest struct {
	UserId    uint64  `json:"userid" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type UpdateUserLocationRequest struct {
	UserId    uint64   `json:"user_id,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

func (req *CreateUserLocationRequest) ToUserLocation() *models.User_location {
	return &models.User_location{
		UserId:    req.UserId,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
