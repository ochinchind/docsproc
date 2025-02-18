package entity

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name      string    `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	Surname   string    `json:"surname"      example:"Doe"                                  gorm:"type:varchar;null;"`
	Email     string    `json:"email"        example:"abc@example.com"                      gorm:"type:varchar;unique;not null"`
	Password  string    `json:"-"            example:"password"                             gorm:"type:varchar;null;"`
	Picture   string    `json:"picture"      example:"https://example.com/picture.jpg"      gorm:"type:varchar;null;"`
	ApiToken  string    `json:"api_token"    example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" gorm:"type:varchar;null;"`
	CreatedAt time.Time `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
}
