package requests

import "github.com/zayn1510/goarchi/app/models"

type LoginlogsRequest struct {
	UserId    uint   `json:"user_id"`
	ISP       string `json:"isp"`
	Username  string `json:"username"`
	Nama      string `json:"nama"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	IPAddress string `json:"ip_address"`
	Device    string `json:"device"`
	Browser   string `json:"browser"`
	Platform  string `json:"platform"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

func (req *LoginlogsRequest) ToModel() *models.LoginLog {
	return &models.LoginLog{
		UserId:    req.UserId,
		ISP:       req.ISP,
		Username:  req.Username,
		Nama:      req.Nama,
		Role:      req.Role,
		Status:    req.Status,
		IPAddress: req.IPAddress,
		Device:    req.Device,
		Browser:   req.Browser,
		Platform:  req.Platform,
		Country:   req.Country,
		City:      req.City,
	}
}
