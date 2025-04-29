package resources

type AuthResource struct {
	Userid   uint   `json:"userid"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}
