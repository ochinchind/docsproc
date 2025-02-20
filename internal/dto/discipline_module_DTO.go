package dto

// UpdateDisciplineModuleDTO -.
type UpdateDisciplineModuleDTO struct {
	Name         string `json:"name" example:"DisciplineModule Name"`
	DisciplineID uint   `json:"discipline_id" example:"1"`
}

// StoreDisciplineModuleDTO -.
type StoreDisciplineModuleDTO struct {
	Name         string `json:"name" example:"DisciplineModule Name" binding:"required"`
	DisciplineID uint   `json:"discipline_id" example:"1" binding:"required"`
}
