package model

type LoginRequest struct {
	Password string `json:"password" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
