package models

// RegisterPasswordRequest register password request dto
type RegisterPasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

// RegisterUsernameRequest register username request dto
type RegisterUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

// RegisterUserRequest register user request dto
type RegisterUserRequest struct {
	Email        string `json:"email" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	SecondName   string `json:"second_name" validate:"required"`
	DepartmentID uint32 `json:"department_id" validate:"required"`
	RoleID       uint32 `json:"role_id" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type RemoveUserRequest struct {
	Id string `json:"id" validate:"required"`
}
