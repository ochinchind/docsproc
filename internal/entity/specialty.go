package entity

import "time"

type Specialty struct {
	ID             uint            `json:"id"             example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name           string          `json:"name"           example:"John"                                 gorm:"type:varchar;not null"`
	Code           string          `json:"code"           example:"JHN"                                  gorm:"type:varchar;not null"`
	CreatedAt      time.Time       `json:"created_at"     example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at"     example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt      time.Time       `json:"-"              example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	Qualifications []Qualification `json:"qualifications"                                                gorm:"foreignKey:SpecialtyID;references:ID"`
}
