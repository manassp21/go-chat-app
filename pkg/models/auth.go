package models

type RegisterRequest struct{
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct{
	Token string `json:"token"`
	User *User `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}