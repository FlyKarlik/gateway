package models

import "time"

type Pattern struct {
	ID string `json:"id"`

	Name string `json:"name"`
	Img  string `json:"img"`
	X    int    `json:"x"`
	Y    int    `json:"y"`

	CreateUserIP string `json:"create_user_ip"`
	UpdateUserIP string `json:"update_user_ip"`
	CreateUserID string `json:"create_user_id"`
	UpdateUserID string `json:"update_user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
