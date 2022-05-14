package models

import (
	"time"

	"gorm.io/gorm"
)

type RecruitmentCycleType string

const (
	PLACEMENT  RecruitmentCycleType = "Placement"
	INTERNSHIP RecruitmentCycleType = "Internship"
)

type RecruitmentCycle struct {
	gorm.Model
	IsActive            bool                 `json:"is_active"`
	AcademicYear        string               `json:"academic_year"`
	Type                RecruitmentCycleType `json:"type"`
	StartDate           time.Time            `json:"start_date"`
	Phase               uint                 `json:"phase"`
	ApplicationCountCap uint                 `json:"application_count_cap"`
}

type RecruitmentCycleQuestionsType string

const (
	MCQ         RecruitmentCycleQuestionsType = "MCQ"
	SHORTANSWER RecruitmentCycleQuestionsType = "ShortAnswer"
	BOOLEAN     RecruitmentCycleQuestionsType = "Boolean"
)

type RecruitmentCycleQuestion struct {
	gorm.Model
	Type               RecruitmentCycleQuestionsType `json:"type"`
	Question           string                        `json:"question"`
	RecruitmentCycleID uint                          `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle              `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Options            string                        `json:"options"` //csv
}

type RecruitmentCycleQuestionsAnswer struct {
	gorm.Model
	RecruitmentCycleQuestionID uint                     `gorm:"index" json:"recruitment_cycle_question_id"`
	RecruitmentCycleQuestion   RecruitmentCycleQuestion `gorm:"foreignkey:RecruitmentCycleQuestionID" json:"-"`
	StudentRecruitmentCycleID  uint                     `gorm:"index" json:"student_recruitment_cycle_id"`
	StudentRecruitmentCycle    StudentRecruitmentCycle  `gorm:"foreignkey:StudentRecruitmentCycleID" json:"-"`
	Answer                     string                   `json:"answer"`
	Comments                   string                   `json:"comments"`
	Status                     string                   `json:"status"`
}

type CompanyRecruitmentCycle struct {
	gorm.Model
	CompanyID          uint             `gorm:"index" json:"company_id"`
	CompanyName        string           `json:"company_name"`
	RecruitmentCycleID uint             `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Comments           string           `json:"comments"`
}

type Notice struct {
	gorm.Model
	CompanyRecruitmentCycleID uint                    `gorm:"index" json:"company_recruitment_cycle_id"`
	CompanyRecruitmentCycle   CompanyRecruitmentCycle `gorm:"foreignkey:CompanyRecruitmentCycleID" json:"-"`
	Title                     string                  `json:"title"`
	Description               string                  `json:"description"`
	CreatedBy                 string                  `json:"created_by"`
	Attachment                string                  `json:"attachment"`
	Tags                      string                  `json:"tags"`
	LastReminderAt            time.Time               `json:"last_reminder_at"`
}

type StudentRecruitmentCycleType string

const (
	PIOPPO    StudentRecruitmentCycleType = "PIO or PPO"
	RECRUITED StudentRecruitmentCycleType = "Recruited"
	AVAILABLE StudentRecruitmentCycleType = "Available"
)

type StudentRecruitmentCycle struct {
	gorm.Model
	StudentID uint                        `gorm:"index" json:"student_id"`
	Email     string                      `grom:"index" json:"email"`
	Type      StudentRecruitmentCycleType `json:"type"`
	IsFrozen  bool                        `json:"is_frozen"`
	Comment   string                      `json:"comment"`
}
