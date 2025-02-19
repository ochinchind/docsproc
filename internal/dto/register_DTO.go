package dto

// RegisterDTO -.
type RegisterDTO struct {
	Username string `json:"username"    example:"john.doe"  binding:"required"`
	Email    string `json:"email"       example:"abc@example.com"  binding:"required,email"`
	Password string `json:"password"    example:"password"  binding:"required"`
	Name     string `json:"name"        example:"John Doe"  binding:"required"`
	Surname  string `json:"surname"     example:"Doe"`
	Phone    string `json:"phone"       example:"+1234567890"`
	Role     string `json:"role"        example:"user"  validate:"oneof=admin user methodologist"`
}
