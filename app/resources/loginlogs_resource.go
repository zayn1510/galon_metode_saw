package resources

import "github.com/zayn1510/goarchi/app/models"

type LoginlogsResource struct {
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
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func NewLoginLogsResource(m models.LoginLog) *LoginlogsResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &LoginlogsResource{ // Mengembalikan pointer agar lebih ringan
		UserId:    m.UserId,
		ISP:       m.ISP,
		Username:  m.Username,
		Nama:      m.Nama,
		Role:      m.Role,
		Status:    m.Status,
		IPAddress: m.IPAddress,
		Device:    m.Device,
		Browser:   m.Browser,
		Platform:  m.Platform,
		Country:   m.Country,
		City:      m.City,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func GetLoginLogsResource(data []models.LoginLog) []*LoginlogsResource {
	resources := make([]*LoginlogsResource, len(data))
	for i, v := range data {
		resources[i] = NewLoginLogsResource(v)
	}
	return resources
}
