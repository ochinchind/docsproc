package entity

import (
	"gorm.io/gorm"
	"time"
)

type DisciplineModuleChapterTopic struct {
	ID                        uint           `json:"id"                   example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name                      string         `json:"name"                 example:"John"                                 gorm:"type:varchar;not null"`
	DisciplineModuleChapterID uint           `json:"discipline_module_chapter_id" example:"1"                                    gorm:"type:bigint;index;not null"`
	HoursTheory               int            `json:"hours_theory"         example:"50"                                   gorm:"type:int;"`
	HoursPractice             int            `json:"hours_practice"       example:"50"                                   gorm:"type:int;"`
	HoursIndividual           int            `json:"hours_individual"     example:"0"                                    gorm:"type:int;"`
	HoursWithTeacher          int            `json:"hours_with_teacher"   example:"0"                                gorm:"type:int;"`
	HoursSelfStudy            int            `json:"hours_self_study"     example:"0"                                gorm:"type:int;"`
	HoursInternship           int            `json:"hours_internship"     example:"0"                                gorm:"type:int;"`
	Type                      string         `json:"type"                 example:"practice"                                gorm:"type:varchar;"`
	CreatedAt                 time.Time      `json:"created_at"           example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt                 time.Time      `json:"updated_at"           example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt                 gorm.DeletedAt `json:"-"                    example:"2021-01-01T00:00:00Z"                 gorm:"index"`
}
