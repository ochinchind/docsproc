package dto

// UpdateDisciplineDTO -.
type UpdateDisciplineDTO struct {
	Name            string `json:"name" example:"Discipline Name"`
	Code            string `json:"code" example:"00342342413"`
	Desc            string `json:"desc"             example:"John Doe"`
	Lang            string `json:"lang"             example:"en"`
	HoursTotal      *int   `json:"hours_total,omitempty"      example:"100"`
	CreaditsCount   *int   `json:"credits_count,omitempty"    example:"100"`
	EducationForm   string `json:"education_form"    example:"full-time"`
	EducationBase   string `json:"education_base"    example:"main"`
	AssessmentType  string `json:"assessment_type"   example:"exam"`
	CompetencyID    uint   `json:"competency_id"     example:"1"`
	QualificationID uint   `json:"qualification_id" example:"1"`
	ContactInfo     string `json:"contact_info"     example:"John Doe"`
}

// StoreDisciplineDTO -.
type StoreDisciplineDTO struct {
	Name            string `json:"name" example:"Discipline Name" binding:"required"`
	Code            string `json:"code" example:"00342342413"`
	Desc            string `json:"desc"             example:"John Doe"`
	Lang            string `json:"lang"             example:"en"`
	HoursTotal      int    `json:"hours_total"      example:"100" binding:"required"`
	CreaditsCount   int    `json:"credits_count"    example:"100" binding:"required"`
	EducationForm   string `json:"education_form"    example:"full-time" binding:"required"`
	EducationBase   string `json:"education_base"    example:"main" binding:"required"`
	AssessmentType  string `json:"assessment_type"   example:"exam"`
	CompetencyID    uint   `json:"competency_id"     example:"1" binding:"required"`
	QualificationID uint   `json:"qualification_id" example:"1" binding:"required"`
	ContactInfo     string `json:"contact_info"     example:"John Doe"`
}
