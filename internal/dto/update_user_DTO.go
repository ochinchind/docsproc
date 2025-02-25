package dto

// UpdateUserDTO -.
type UpdateUserDTO struct {
	Username string `json:"username"    example:"john.doe"`
	Email    string `json:"email"       example:"abc@example.com" validate:"email"`
	Name     string `json:"name"        example:"John Doe"`
	Password string `json:"password"    example:"password"`
	Surname  string `json:"surname"     example:"Doe"`
	Phone    string `json:"phone"       example:"+1234567890"`
	Role     string `json:"role"        example:"user"  validate:"oneof=admin teacher methodologist"`
}
