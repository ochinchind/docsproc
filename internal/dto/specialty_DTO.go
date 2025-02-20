package dto

// UpdateSpecialtyDTO -.
type UpdateSpecialtyDTO struct {
	Name string `json:"name" example:"Specialty Name" validate:"required"`
	Code string `json:"code" example:"00342342413" validate:"required"`
}
