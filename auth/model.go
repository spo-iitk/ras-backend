package auth

import (
	"github.com/spo-iitk/ras-backend/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       string         `gorm:"uniqueIndex" json:"user_id"`
	Password     string         `json:"password"`
	RoleID       constants.Role `json:"role_id" gorm:"default:1"` // student role by default
	Name         string         `json:"name"`
	IsActive     bool           `json:"is_active" gorm:"default:true"`
	LastLogin    uint           `json:"last_login" gorm:"index;autoUpdateTime:milli"`
	RefreshToken string         `json:"refresh_token"`
}

type OTP struct {
	gorm.Model
	UserID  string `gorm:"column:user_id"`
	OTP     string `gorm:"column:otp"`
	Expires uint   `gorm:"column:expires"`
}

type CompanySignUpRequest struct {
	gorm.Model
	CompanyName string `json:"company_name" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Designation string `json:"designation" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	IsReviewed  bool   `json:"is_reviewed"`
	Comments    string `json:"comments"`
}
