package auth

import (
	"gorm.io/gorm"
)

type Role uint8

const (
	GOD     Role = 0
	OPC     Role = 1
	APC     Role = 2
	DPC     Role = 3
	COCO    Role = 4
	STAFF   Role = 5
	COMPANY Role = 11
	STUDENT Role = 10
)

type User struct {
	gorm.Model
	UserID       string `gorm:"uniqueIndex" json:"user_id"`
	Password     string `json:"password"`
	RoleID       Role   `json:"role_id" gorm:"default:10"` // student role by default
	Name         string `json:"name"`
	IsActive     bool   `json:"is_active" gorm:"default:true"`
	LastLogin    uint   `json:"last_login" gorm:"index;autoUpdateTime:milli"`
	RefreshToken string `json:"refresh_token"`
}

type OTP struct {
	gorm.Model
	UserID  string `gorm:"column:user_id"`
	OTP     string `gorm:"column:otp"`
	Expires uint   `gorm:"column:expires"`
}

type CompanySignUpRequest struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
