package request

// LoginRequest struct
type LoginRequest struct {
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
}
