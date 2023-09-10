package model

import (
	"time"
)

type User struct {
	Id          int       `gorm:"primaryKey;auto_increment" json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Roles       []Role    `gorm:"many2many:user_role;" json:"roles"`
}

type CreateUser struct {
	UserName    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func (mc *User) TableName() string {
	return "user"
}
