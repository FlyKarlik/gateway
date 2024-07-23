package models

// CheckUsernameRequest check username dto
type CheckUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}
