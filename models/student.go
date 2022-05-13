package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	RollNo string `gorm:"uniqueIndex" json:"roll_no"`
	Name   string `json:"name"`
	//!TODO: see alternative fuck legacy
	ProgramID             uint       `gorm:"column:program_id" json:"program_id"`
	Program               Program    `gorm:"foreignkey:ProgramID" json:"-"`
	DepartmentID          uint       `gorm:"column:department_id" json:"department_id"`
	Department            Department `gorm:"foreignkey:DepartmentID" json:"-"`
	SecondaryDepartmentID uint       `gorm:"column:secondary_department_id" json:"secondary_department_id"`
	SecondaryDepartment   Department `gorm:"foreignkey:SecondaryDepartmentID" json:"-"`
	//!TODO: What is this?
	Specialization         string    `json:"specialization"`
	Preference             string    `json:"preference"`
	Gender                 string    `json:"gender"`
	Disablity              string    `json:"disability"`
	DOB                    time.Time `json:"dob"`
	ExpectedGraduationYear uint      `json:"expected_graduation_year"`
	IITKEmail              string    `gorm:"uniqueIndex" json:"iitk_email"`
	PersonalEmail          string    `json:"personal_email"`
	Phone                  string    `json:"phone"`
	AlternatePhone         string    `json:"alternate_phone"`
	WhatsappNumber         string    `json:"whatsapp_number"`
	CurrentCPI             float64   `json:"current_cpi"`
	UGCPI                  float64   `json:"ug_cpi"`
	TenthBoard             string    `json:"tenth_board"`
	TenthYear              uint      `json:"tenth_year"`
	TenthMarks             float64   `json:"tenth_marks"`
	TwelfthBoard           string    `json:"twelfth_board"`
	TwelfthYear            uint      `json:"twelfth_year"`
	TwelfthMarks           float64   `json:"twelfth_marks"`
	EntranceExam           string    `json:"entrance_exam"`
	EntranceExamRank       uint      `json:"entrance_exam_rank"`
	Category               string    `json:"category"`
	CategoryRank           uint      `json:"category_rank"`
	CurrentAddress         string    `gorm:"type:longtext" json:"current_address"`
	PermanentAddress       string    `gorm:"type:longtext" json:"permanent_address"`
	FriendName             string    `json:"friend_name"`
	FriendPhone            string    `json:"friend_phone"`
	IsEditable             bool      `json:"is_editable"`
}

type Program struct {
	gorm.Model
	Name string `json:"name"`
}

type Department struct {
	gorm.Model
	Name string `json:"name"`
}

type ProgramDepartment struct {
	gorm.Model
	ProgramID    uint       `gorm:"column:program_id" json:"program_id"`
	Program      Program    `gorm:"foreignkey:ProgramID" json:"-"`
	DepartmentID uint       `gorm:"column:department_id" json:"department_id"`
	Department   Department `gorm:"foreignkey:DepartmentID" json:"-"`
}
