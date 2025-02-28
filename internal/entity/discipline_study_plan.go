package entity

import (
	"gorm.io/gorm"
	"time"
)

// DisciplineStudyPlan -.
type DisciplineStudyPlan struct {
	ID             uint           `json:"id"               example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	DisciplineID   uint           `json:"discipline_id"    example:"1"                                    gorm:"type:bigint;index;not null"`
	PreRequisites  string         `json:"pre_requisites"   example:"1,2,3"                               gorm:"type:text;"`
	PostRequisites string         `json:"post_requisites"  example:"1,2,3"                              gorm:"type:text;"`
	Necessities    string         `json:"necessities"      example:"1,2,3"                               gorm:"type:text;"`
	ContactInfo    string         `json:"contact_info"     example:"John Doe"                             gorm:"type:text;"`
	CreatedAt      time.Time      `json:"created_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at"       example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-"            example:"2021-01-01T00:00:00Z"                     gorm:"index"`
	Discipline     *Discipline    `json:"discipline"                                  gorm:"foreignKey:DisciplineID"`
}
