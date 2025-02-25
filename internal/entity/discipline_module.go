package entity

import (
	"gorm.io/gorm"
	"time"
)

type DisciplineModule struct {
	ID                       uint                      `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name                     string                    `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	FirstSemester            int                       `json:"first_semester" example:"73"                                  gorm:"type:int;not null;default:0"`
	SecondSemester           int                       `json:"second_semester" example:"48"                                 gorm:"type:int;not null;default:0"`
	ThirdSemester            int                       `json:"third_semester" example:"0"                                  gorm:"type:int;not null;default:0"`
	FourthSemester           int                       `json:"fourth_semester" example:"0"                                 gorm:"type:int;not null;default:0"`
	FifthSemester            int                       `json:"fifth_semester" example:"1"                                  gorm:"type:int;not null;default:0"`
	SixthSemester            int                       `json:"sixth_semester" example:"1"                                  gorm:"type:int;not null;default:0"`
	SeventhSemester          int                       `json:"seventh_semester" example:"1"                                 gorm:"type:int;not null;default:0"`
	EighthSemester           int                       `json:"eighth_semester" example:"1"                                 gorm:"type:int;not null;default:0"`
	DisciplineID             uint                      `json:"discipline_id" example:"1"                                    gorm:"type:bigint;index;not null"`
	CreatedAt                time.Time                 `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt                time.Time                 `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt                gorm.DeletedAt            `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
	DisciplineModuleChapters []DisciplineModuleChapter `json:"discipline_module_chapters"                                  gorm:"foreignKey:DisciplineModuleID;references:ID"`
}
