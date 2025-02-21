package dto

// UpdateDisciplineDTO -.
type UpdateDisciplineDTO struct {
	Name             string `json:"name" example:"Discipline Name"`
	Code             string `json:"code" example:"00342342413"`
	Desc             string `json:"desc"             example:"John Doe"  `
	Lang             string `json:"lang"             example:"en"        `
	HoursTotal       *int   `json:"hours_total,omitempty"      example:"100"`
	HoursTheory      *int   `json:"hours_theory,omitempty"     example:"50" `
	HoursPractice    *int   `json:"hours_practice,omitempty"   example:"50" `
	HoursIndividual  *int   `json:"hours_individual,omitempty" example:"0"  `
	HoursWithTeacher *int   `json:"hours_with_teacher,omitempty" example:"0"`
	HoursSelfStudy   *int   `json:"hours_self_study,omitempty"  example:"0" `
	HoursInternship  *int   `json:"hours_internship,omitempty"  example:"0" `
	AssessmentType   string `json:"assessment_type"   example:"exam"     `
	Competencies     string `json:"competencies"     example:"1,2,3"     `
	PreRequisites    string `json:"pre_requisites"   example:"1,2,3"     `
	PostRequisites   string `json:"post_requisites"  example:"1,2,3"     `
	Necessities      string `json:"necessities"      example:"1,2,3"     `
	QualificationID  uint   `json:"qualification_id" example:"1"`
}

// StoreDisciplineDTO -.
type StoreDisciplineDTO struct {
	Name             string `json:"name" example:"Discipline Name" binding:"required"`
	Code             string `json:"code" example:"00342342413"`
	Desc             string `json:"desc"             example:"John Doe"`
	Lang             string `json:"lang"             example:"en"   `
	HoursTotal       int    `json:"hours_total"      example:"100"  `
	HoursTheory      int    `json:"hours_theory"     example:"50"   `
	HoursPractice    int    `json:"hours_practice"   example:"50"   `
	HoursIndividual  int    `json:"hours_individual" example:"0"    `
	HoursWithTeacher int    `json:"hours_with_teacher" example:"0"  `
	HoursSelfStudy   int    `json:"hours_self_study"  example:"0"   `
	HoursInternship  int    `json:"hours_internship"  example:"0"   `
	AssessmentType   string `json:"assessment_type"   example:"exam"`
	Competencies     string `json:"competencies"     example:"1,2,3"`
	PreRequisites    string `json:"pre_requisites"   example:"1,2,3"`
	PostRequisites   string `json:"post_requisites"  example:"1,2,3"`
	Necessities      string `json:"necessities"      example:"1,2,3"`
	QualificationID  uint   `json:"qualification_id" example:"1" binding:"required"`
}
