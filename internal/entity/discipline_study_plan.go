package entity

// DisciplineStudyPlan -.
type DisciplineStudyPlan struct {
	ID             uint   `json:"id"               example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	DisciplineID   uint   `json:"discipline_id"    example:"1"                                    gorm:"type:bigint;index;not null"`
	PreRequisites  string `json:"pre_requisites"   example:"1,2,3"                               gorm:"type:text;"`
	PostRequisites string `json:"post_requisites"  example:"1,2,3"                              gorm:"type:text;"`
	Necessities    string `json:"necessities"      example:"1,2,3"                               gorm:"type:text;"`
	ContactInfo    string `json:"contact_info"     example:"John Doe"                             gorm:"type:text;"`
}
