package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       string    `gorm:"uniqueIndex" json:"user_id"`
	Password     string    `json:"password"`
	RoleID       uint      `gorm:"column:role_id" json:"role_id"`
	Role         Role      `gorm:"foreignkey:RoleID" json:"-"`
	Name         string    `json:"name"`
	IsActive     bool      `json:"is_active"`
	LastLogin    time.Time `json:"last_login"`
	RefreshToken string    `json:"refresh_token"`
}

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type Master struct {
	gorm.Model
	UserID   string `gorm:"column:user_id" json:"user_id"`
	User     User   `gorm:"foreignkey:UserID" json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
