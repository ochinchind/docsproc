package dto

// UpdateDisciplineModuleTopicDTO -.
type UpdateDisciplineModuleTopicDTO struct {
	Name               string `json:"name" example:"DisciplineModuleTopic Name"`
	HoursTheory        int    `json:"hours_theory"     example:"50" `
	HoursPractice      int    `json:"hours_practice"   example:"50" `
	HoursIndividual    int    `json:"hours_individual" example:"0"  `
	HoursWithTeacher   int    `json:"hours_with_teacher" example:"0"`
	HoursSelfStudy     int    `json:"hours_self_study"  example:"0" `
	HoursInternship    int    `json:"hours_internship"  example:"0" `
	Type               string `json:"type"                 example:"practice"   binding:"omitempty,oneof=practice theory"`
	DisciplineModuleID uint   `json:"discipline_module_id" example:"1"`
}

// StoreDisciplineModuleTopicDTO -.
type StoreDisciplineModuleTopicDTO struct {
	Name               string `json:"name" example:"DisciplineModuleTopic Name" binding:"required"`
	HoursTheory        int    `json:"hours_theory"     example:"50" `
	HoursPractice      int    `json:"hours_practice"   example:"50" `
	HoursIndividual    int    `json:"hours_individual" example:"0"  `
	HoursWithTeacher   int    `json:"hours_with_teacher" example:"0"`
	HoursSelfStudy     int    `json:"hours_self_study"  example:"0" `
	HoursInternship    int    `json:"hours_internship"  example:"0" `
	Type               string `json:"type"                 example:"practice"   binding:"required,oneof=practice theory"`
	DisciplineModuleID uint   `json:"discipline_module_id" example:"1" binding:"required"`
}
