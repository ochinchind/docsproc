package dto

// UpdateQualificationDTO -.
type UpdateQualificationDTO struct {
	Name        string `json:"name" example:"Qualification Name"`
	Code        string `json:"code" example:"00342342413"`
	SpecialtyID uint   `json:"specialty_id" example:"1"`
}

// StoreQualificationDTO -.
type StoreQualificationDTO struct {
	Name        string `json:"name" example:"Qualification Name" binding:"required"`
	Code        string `json:"code" example:"00342342413" binding:"required"`
	SpecialtyID uint   `json:"specialty_id" example:"1" binding:"required"`
}
