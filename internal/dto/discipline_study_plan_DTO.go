package dto

// UpdateDisciplineStudyPlanDTO -.
type UpdateDisciplineStudyPlanDTO struct {
	DisciplineID   uint   `json:"discipline_id"    example:"1"`
	PreRequisites  string `json:"pre_requisites"   example:"1,2,3"`
	PostRequisites string `json:"post_requisites"  example:"1,2,3"`
	Necessities    string `json:"necessities"      example:"1,2,3"`
	ContactInfo    string `json:"contact_info"     example:"John Doe"`
}

// StoreDisciplineStudyPlanDTO -.
type StoreDisciplineStudyPlanDTO struct {
	DisciplineID   uint   `json:"discipline_id"    example:"1"`
	PreRequisites  string `json:"pre_requisites"   example:"1,2,3"`
	PostRequisites string `json:"post_requisites"  example:"1,2,3"`
	Necessities    string `json:"necessities"      example:"1,2,3"`
	ContactInfo    string `json:"contact_info"     example:"John Doe"`
}
