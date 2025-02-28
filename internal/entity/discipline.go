package entity

import (
	"gorm.io/gorm"
	"time"
)

type Discipline struct {
	ID                  uint                  `json:"id"               example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name                string                `json:"name"             example:"John"                                 gorm:"type:varchar;not null"`
	Code                string                `json:"code"             example:"JHN"                                  gorm:"type:varchar;"`
	Desc                string                `json:"desc"             example:"John Doe"                             gorm:"type:varchar;"`
	Lang                string                `json:"lang"             example:"en"                                   gorm:"type:varchar;"`
	HoursTotal          int                   `json:"hours_total"      example:"100"                                  gorm:"type:int;not null"`
	CreaditsCount       int                   `json:"credits_count"    example:"100"                                  gorm:"type:int;not null"`
	EducationForm       string                `json:"education_form"    example:"full-time"                          gorm:"type:varchar;not null"`
	EducationBase       string                `json:"education_base"    example:"main"                               gorm:"type:varchar;not null"`
	AssessmentType      string                `json:"assessment_type"   example:"exam"                                gorm:"type:varchar;"`
	CompetencyID        uint                  `json:"competency_id"     example:"1"                                   gorm:"type:bigint;index;not null"`
	QualificationID     uint                  `json:"qualification_id" example:"1"                                   gorm:"type:bigint;index;not null"`
	CreatedAt           time.Time             `json:"created_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt           time.Time             `json:"updated_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt        `json:"-"            example:"2021-01-01T00:00:00Z"                     gorm:"index"`
	DisciplineModules   []DisciplineModule    `json:"discipline_modules"                                                 gorm:"foreignKey:DisciplineID;references:ID"`
	DisciplineStudyPlan []DisciplineStudyPlan `json:"discipline_study_plans"                                              gorm:"foreignKey:DisciplineID;references:ID"`
}
