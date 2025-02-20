package entity

import (
	"gorm.io/gorm"
	"time"
)

type Specialty struct {
	ID             uint            `json:"id"             example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name           string          `json:"name"           example:"John"                                 gorm:"type:varchar;not null"                binding:"required"`
	Code           string          `json:"code"           example:"JHN"                                  gorm:"type:varchar;not null"                binding:"required"`
	CreatedAt      time.Time       `json:"created_at"     example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at"     example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `json:"-"              example:"2021-01-01T00:00:00Z"                 gorm:"index"`
	Qualifications []Qualification `json:"qualifications"                                                gorm:"foreignKey:SpecialtyID;references:ID"`
}
