package models

import "time"

type Layer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LayerType    string `json:"layer_type"`
	TableID      string `json:"table_id"`
	CreateUserID string `json:"create_user_id"`
	CreateUserIP string `json:"create_user_ip"`
	UpdateUserID string `json:"update_user_id"`
	UpdateUserIP string `json:"update_user_ip"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// ErrorMessage ...
type ErrorMessage struct {
	MessageID string    `json:"messageId"`
	Offset    int64     `json:"offset"`
	Partition int       `json:"partition"`
	Topic     string    `json:"topic"`
	Error     string    `json:"error"`
	Time      time.Time `json:"time"`
}
