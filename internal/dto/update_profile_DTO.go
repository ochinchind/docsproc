package dto

// UpdateProfileDTO -.
type UpdateProfileDTO struct {
	Username string `json:"username"    example:"john.doe"`
	Name     string `json:"name"        example:"John Doe"`
	Password string `json:"password"    example:"password"`
	Surname  string `json:"surname"     example:"Doe"`
	Phone    string `json:"phone"       example:"+1234567890"`
}
