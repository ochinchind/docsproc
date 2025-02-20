package dto

// UpdateSpecialtyDTO -.
type UpdateSpecialtyDTO struct {
	Name string `json:"name" example:"Specialty Name"`
	Code string `json:"code" example:"00342342413"`
}

// StoreSpecialtyDTO -.
type StoreSpecialtyDTO struct {
	Name string `json:"name" example:"Specialty Name" binding:"required"`
	Code string `json:"code" example:"00342342413" binding:"required"`
}
