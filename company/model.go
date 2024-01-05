package company

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Tags        string `json:"tags"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

type CompanyHR struct {
	gorm.Model
	CompanyID   uint    `json:"company_id"`
	Company     Company `gorm:"foreignkey:CompanyID" json:"-"`
	Name        string  `json:"name"`
	Email       string  `gorm:"uniqueIndex;->;<-:create" json:"email"`
	Phone       string  `json:"phone"`
	Designation string  `json:"designation"`
}
