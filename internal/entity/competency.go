package entity

import (
	"gorm.io/gorm"
	"time"
)

// Competency -.
type Competency struct {
	ID        uint           `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name      string         `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	CreatedAt time.Time      `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
}
