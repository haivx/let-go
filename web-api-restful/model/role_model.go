package model

const (
	ADMIN = "administration"
	USER  = "user"
	MOD   = "moderator"
)

type Role struct {
	Id          int          `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Name        string       `json:"name" validate:"ADMIN|USER|MOD"`
	Permissions []Permission `json:"-" gorm:"many2many:role_permission;"`
	Users       []User       `json:"-" gorm:"many2many:user_role;"`
}

type UserRole struct {
	UserId int `gorm:"primaryKey" column:"user_id"`
	RoleId int `gorm:"primaryKey" column:"role_id"`
}

func (r *Role) TableName() string {
	return "role"
}

func (m *UserRole) TableName() string {
	return "user_role"
}
