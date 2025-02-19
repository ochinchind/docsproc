package entity

import (
	"gorm.io/gorm"
	"time"
)

type Discipline struct {
	ID                uint               `json:"id"               example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name              string             `json:"name"             example:"John"                                 gorm:"type:varchar;not null"`
	Code              string             `json:"code"             example:"JHN"                                  gorm:"type:varchar;"`
	Desc              string             `json:"desc"             example:"John Doe"                             gorm:"type:varchar;"`
	Lang              string             `json:"lang"             example:"en"                                   gorm:"type:varchar;"`
	HoursTotal        int                `json:"hours_total"      example:"100"                                  gorm:"type:int;"`
	HoursTheory       int                `json:"hours_theory"     example:"50"                                   gorm:"type:int;"`
	HoursPractice     int                `json:"hours_practice"   example:"50"                                   gorm:"type:int;"`
	HoursIndividual   int                `json:"hours_individual" example:"0"                                    gorm:"type:int;"`
	HoursWithTeacher  int                `json:"hours_with_teacher" example:"0"                                gorm:"type:int;"`
	HoursSelfStudy    int                `json:"hours_self_study"  example:"0"                                gorm:"type:int;"`
	HoursInternship   int                `json:"hours_internship"  example:"0"                                gorm:"type:int;"`
	AssessmentType    string             `json:"assessment_type"   example:"exam"                                gorm:"type:varchar;"`
	Competencies      string             `json:"competencies"     example:"1,2,3"                               gorm:"type:text;"`
	PreRequisites     string             `json:"pre_requisites"   example:"1,2,3"                               gorm:"type:text;"`
	PostRequisites    string             `json:"post_requisites"  example:"1,2,3"                              gorm:"type:text;"`
	Necessities       string             `json:"necessities"      example:"1,2,3"                               gorm:"type:text;"`
	QualificationID   uint               `json:"qualification_id" example:"1"                                gorm:"type:bigint;index;not null"`
	CreatedAt         time.Time          `json:"created_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt         time.Time          `json:"updated_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt     `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
	DisciplineModules []DisciplineModule `json:"discipline_modules"                                                 gorm:"foreignKey:DisciplineID;references:ID"`
}
