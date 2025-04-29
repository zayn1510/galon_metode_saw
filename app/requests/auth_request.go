package requests

type AuthRequest struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
type UpdatePasswordRequest struct {
	Username           string `form:"username" json:"username" validate:"required"`
	Password           string `form:"password" json:"password" validate:"required"`
	NewPassword        string `form:"new_password" json:"new_password" validate:"required"`
	ConfirmNewPassword string `form:"confirm_new_password" json:"confirm_new_password" validate:"required"`
}
