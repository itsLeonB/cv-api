package model

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=89"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}
