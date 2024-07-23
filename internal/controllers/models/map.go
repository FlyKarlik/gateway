package models

import "time"

type Map struct {
	ID string `gorm:"primaryKey;type:uuid" json:"id"`

	Name     string `gorm:"not null" json:"name"`
	Picture  string `gorm:"type:text;not null" json:"picture"`
	Describe string `gorm:"type:text;not null" json:"describe"`
	Active   bool   `gorm:"not null;default:true" json:"active"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
