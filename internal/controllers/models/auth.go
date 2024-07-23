package models

// AuthorizeRequest authorize request dto
type AuthorizeRequest struct {
	Email string `json:"email" validate:"required"`
}

// VerifyAuthorizeRequest verify authorize request dto
type VerifyAuthorizeRequest struct {
	AuthorizeToken string `json:"authorize_token" validate:"required"`
	Code           string `json:"code" validate:"required"`
}

// FirstLoginRequest first login request dto
type FirstLoginRequest struct {
	Password              string `json:"password" validate:"required"`
	DeviceName            string `json:"device_name"`
	UnrestrictedPublicKey string `json:"unrestricted_public_key"`
}

// VerifyFirstLoginRequest verify first login request dto
type VerifyFirstLoginRequest struct {
	ChallengeID                  string `json:"challenge_id" validate:"required"`
	DeviceID                     string `json:"device_id" validate:"required"`
	OTPSignature                 string `json:"otp_signature" validate:"required"`
	RestrictedPublicKey          string `json:"restricted_public_key" validate:"required"`
	RestrictedPublicKeySignature string `json:"restricted_public_key_signature" validate:"required"`
}

// RecurringLoginRequest recurring login request dto
type RecurringLoginRequest struct {
	Password    string `json:"password"`
	DeviceToken string `json:"device_token"`
}

// LoginUserRequest login user
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
