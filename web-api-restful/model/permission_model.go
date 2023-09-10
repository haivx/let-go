package model

import (
	"time"
)

type Permission struct {
	Id         int       `gorm:"primaryKey;auto_increment" json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Created_at time.Time `json:"created_at"`
	Roles      []Role    `gorm:"many2many:role_permission;"`
}

type RolePermission struct {
	PermissionId int `gorm:"primaryKey" column:"permission_id"`
	RoleId       int `gorm:"primaryKey" column:"role_id"`
}

func (m *Permission) TableName() string {
	return "permission"
}

func (m *RolePermission) TableName() string {
	return "role_permission"
}
