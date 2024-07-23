package models

// RefreshTokenRequest refresh token dto
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
