package dto

type CreateRole struct {
	Name string `json:"name" validate:"oneof=ADMIN USER MOD"`
}
