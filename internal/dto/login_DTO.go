package dto

// LoginDTO -.
type LoginDTO struct {
	Username string `json:"username"    example:"john.doe"  binding:"required"`
	Password string `json:"password"    example:"password"  binding:"required"`
}
