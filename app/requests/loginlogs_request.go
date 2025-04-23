package requests

import "github.com/zayn1510/goarchi/app/models"

type LoginlogsRequest struct {
	UserId    uint   `json:"user_id"`
	ISP       string `json:"isp,omitempty"`
	Username  string `json:"username,omitempty"`
	Nama      string `json:"nama,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"status,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	Device    string `json:"device,omitempty"`
	Browser   string `json:"browser,omitempty"`
	Platform  string `json:"platform,omitempty"`
	Country   string `json:"country,omitempty"`
	City      string `json:"city,omitempty"`
	Token     string `json:"token,omitempty"`
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
