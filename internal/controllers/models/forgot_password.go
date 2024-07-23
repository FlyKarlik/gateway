package models

// VerifyForgotPasswordRequest verify forgot password dto
type VerifyForgotPasswordRequest struct {
	ForgotToken string `json:"forgot_token" validate:"required"`
	Code        string `json:"code" validate:"required"`
}

// ResetPasswordRequest reset password dto
type ResetPasswordRequest struct {
	ResetPasswordToken string `json:"reset_password_token" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
}
