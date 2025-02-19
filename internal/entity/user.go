package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id"           example:"1"                                    gorm:"primaryKey:autoIncrement;"`
	Name      string         `json:"name"         example:"John"                                 gorm:"type:varchar;not null"`
	Surname   string         `json:"surname"      example:"Doe"                                  gorm:"type:varchar;null;"`
	Phone     string         `json:"phone"        example:"877755544434"                         gorm:"type:varchar;null;"`
	Username  string         `json:"username"     example:"johndoe"                              gorm:"type:varchar;unique;not null"`
	Email     string         `json:"email"        example:"abc@example.com"                      gorm:"type:varchar;unique;not null"`
	Password  string         `json:"-"            example:"password"                             gorm:"type:varchar;null;"`
	Picture   string         `json:"picture"      example:"https://example.com/picture.jpg"      gorm:"type:varchar;null;"`
	ApiToken  string         `json:"api_token"    example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" gorm:"type:varchar;null;"`
	Role      string         `json:"role"         example:"admin"                                gorm:"type:varchar;not null;default:'user'"`
	CreatedAt time.Time      `json:"created_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at"   example:"2021-01-01T00:00:00Z"                 gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"            example:"2021-01-01T00:00:00Z"                 gorm:"index"`
}
