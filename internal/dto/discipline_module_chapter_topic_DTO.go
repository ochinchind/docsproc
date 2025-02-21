package dto

// UpdateDisciplineModuleChapterTopicDTO -.
type UpdateDisciplineModuleChapterTopicDTO struct {
	Name                      string `json:"name" example:"DisciplineModuleChapterTopic Name"`
	HoursTheory               *int   `json:"hours_theory,omitempty" example:"50"`
	HoursPractice             *int   `json:"hours_practice,omitempty" example:"50"`
	HoursIndividual           *int   `json:"hours_individual,omitempty" example:"0"`
	HoursWithTeacher          *int   `json:"hours_with_teacher,omitempty" example:"0"`
	HoursSelfStudy            *int   `json:"hours_self_study,omitempty" example:"0"`
	HoursInternship           *int   `json:"hours_internship,omitempty" example:"0"`
	Type                      string `json:"type" example:"practice" binding:"omitempty,oneof=practice theory combined"`
	DisciplineModuleChapterID uint   `json:"discipline_module_chapter_id" example:"1"`
}

// StoreDisciplineModuleChapterTopicDTO -.
type StoreDisciplineModuleChapterTopicDTO struct {
	Name                      string `json:"name" example:"DisciplineModuleChapterTopic Name" binding:"required"`
	HoursTheory               int    `json:"hours_theory,omitempty"     example:"50" `
	HoursPractice             int    `json:"hours_practice,omitempty"   example:"50" `
	HoursIndividual           int    `json:"hours_individual,omitempty" example:"0"  `
	HoursWithTeacher          int    `json:"hours_with_teacher,omitempty" example:"0"`
	HoursSelfStudy            int    `json:"hours_self_study,omitempty"  example:"0" `
	HoursInternship           int    `json:"hours_internship,omitempty"  example:"0" `
	Type                      string `json:"type"                 example:"practice"   binding:"required,oneof=practice theory combined"`
	DisciplineModuleChapterID uint   `json:"discipline_module_chapter_id" example:"1" binding:"required"`
}
