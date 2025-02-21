package dto

// UpdateDisciplineModuleChapterDTO -.
type UpdateDisciplineModuleChapterDTO struct {
	Name               string `json:"name" example:"DisciplineModuleChapter Name"`
	DisciplineModuleID uint   `json:"discipline_module_id" example:"1"`
}

// StoreDisciplineModuleChapterDTO -.
type StoreDisciplineModuleChapterDTO struct {
	Name               string `json:"name" example:"DisciplineModuleChapter Name" binding:"required"`
	DisciplineModuleID uint   `json:"discipline_module_id" example:"1" binding:"required"`
}
