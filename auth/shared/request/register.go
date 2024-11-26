package request

type RegisterRequest struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	TelegramLink   string `json:"telegram_link,omitempty"`
}
