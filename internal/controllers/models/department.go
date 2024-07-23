package models

// AddUserDepartmentRequest struct for request addUserDepartment
type AddUserDepartmentRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

// RemoveUserDepartmentRequest struct for request removeUserDepartment
type RemoveUserDepartmentRequest struct {
	Id uint32 `json:"id" validate:"required"`
}

// GetUserDepartmentRequest struct for getUserDepartment
type GetUserDepartmentRequest struct {
	Id uint32 `json:"id" validate:"required"`
}
