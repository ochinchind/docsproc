package dto

// UpdateCompetencyDTO -.
type UpdateCompetencyDTO struct {
	Name string `json:"name" example:"Competency Name"`
}

// StoreCompetencyDTO -.
type StoreCompetencyDTO struct {
	Name string `json:"name" example:"Competency Name" binding:"required"`
}
