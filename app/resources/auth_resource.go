package resources

type AuthResource struct {
	ID        int    `json:"id"`
	ISP       string `json:"isp"`
	Username  string `json:"username"`
	Nama      string `json:"nama"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	LoginTime string `json:"login_time"`
	IPAddress string `json:"ip_address"`
	Device    string `json:"device"`
	Browser   string `json:"browser"`
	Platform  string `json:"platform"`
	Country   string `json:"country"`
	City      string `json:"city"`
}
