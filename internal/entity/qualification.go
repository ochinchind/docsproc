package entity

import (
	"gorm.io/gorm"
	"time"
)

type Qualification struct {
	ID          uint           `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name        string         `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	Code        string         `json:"code"         example:"JHN"                                  gorm:"type:varchar;not null"`
	SpecialtyID uint           `json:"specialty_id" example:"1"                                    gorm:"type:bigint;index;not null"`
	CreatedAt   time.Time      `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
	Disciplines []Discipline   `json:"disciplines"                                                 gorm:"foreignKey:QualificationID;references:ID"`
}
