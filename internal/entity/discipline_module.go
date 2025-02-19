package entity

import (
	"gorm.io/gorm"
	"time"
)

type DisciplineModule struct {
	ID                     uint                    `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name                   string                  `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	DisciplineID           uint                    `json:"discipline_id" example:"1"                                    gorm:"type:bigint;index;not null"`
	CreatedAt              time.Time               `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt              time.Time               `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt              gorm.DeletedAt          `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
	DisciplineModuleTopics []DisciplineModuleTopic `json:"discipline_module_topics"                                                 gorm:"foreignKey:DisciplineModuleID;references:ID"`
}
