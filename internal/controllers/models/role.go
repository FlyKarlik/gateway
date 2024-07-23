package models

// AddUserRoleRequest struct for request addUserRole
type AddUserRoleRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

// RemoveUserRoleRequest struct for request removeUserRole
type RemoveUserRoleRequest struct {
	Id uint32 `json:"id" validate:"required"`
}

// GetUserRoleRequest get role by id struct
type GetUserRoleRequest struct {
	Id uint32 `json:"id" validate:"required"`
}
