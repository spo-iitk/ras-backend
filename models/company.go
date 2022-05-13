package models

import "gorm.io/gorm"

type CompanySignup struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Designation string `json:"designation"`
	Phone       string `json:"phone"`
	IsReviewed  bool   `json:"is_reviewed"`
	Comments    string `json:"comments"`
}

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
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Designation string  `json:"designation"`
}
