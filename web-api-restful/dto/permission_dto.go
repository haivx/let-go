package dto

type CreatePermission struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type CreateRolePermission struct {
	PermissionId int `json:"permission_id" validate:"required"`
	RoleId       int `json:"role_id" validate:"required"`
}
