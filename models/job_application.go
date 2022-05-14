package models

import (
	"time"

	"gorm.io/gorm"
)

type JobProforma struct {
	gorm.Model
	CompanyID                 uint      `gorm:"index" json:"company_id"`
	CompanyRecruitmentCycleID uint      `gorm:"index" json:"company_recruitment_cycle_id"`
	RecruitmentCycleID        uint      `gorm:"index" json:"recruitment_cycle_id"`
	IsApproved                bool      `json:"is_approved"`
	ActionTakenBy             string    `json:"action_taken_by"`
	SetDeadline               time.Time `json:"set_deadline"`
	HideDetails               bool      `json:"hide_details"` //gorm default value is false
	ActiveHRID                string    `json:"active_hr_id"`
	NatureOfBusiness          string    `json:"nature_of_business"`
	TentativeJobLocation      string    `json:"tentative_job_location"`
	JobDescription            string    `json:"job_description"`
	CostToCompany             string    `json:"cost_to_company"`
	PackageDetails            string    `json:"package_details"`
	BondDetails               string    `json:"bond_details"`
	MedicalRequirements       string    `json:"medical_requirements"`
	AdditionalEligibility     string    `json:"additional_eligibility"`
	MessageForCordinator      string    `json:"message_for_cordinator"`
}

// ELIGIBILITY MATRIX

type JobApplicationQuestion struct {
	gorm.Model
	Type          RecruitmentCycleQuestionsType `json:"type"`
	Question      string                        `json:"question"`
	JobPerformaID uint                          `gorm:"index" json:"job_performa_id"`
	JobPerforma   JobProforma                   `gorm:"foreignkey:JobPerformaID" json:"-"`
	Options       string                        `json:"options"` //csv
}

type JobApplicationQuestionsAnswer struct {
	gorm.Model
	JobApplicationQuestionID  uint                   `gorm:"index" json:"job_application_question_id"`
	JobApplicationQuestion    JobApplicationQuestion `gorm:"foreignkey:JobApplicationQuestionID" json:"-"`
	StudentRecruitmentCycleID uint                   `gorm:"index" json:"student_recruitment_cycle_id"`
	Answer                    string                 `json:"answer"`
}

type JobPerformaEvent struct {
	gorm.Model
	JobPerformaID    uint        `gorm:"index" json:"job_performa_id"`
	JobPerforma      JobProforma `gorm:"foreignkey:JobPerformaID" json:"-"`
	Name             string      `json:"name"`
	Duration         string      `json:"duration"`
	Venue            string      `json:"venue"`
	StartTime        time.Time   `json:"start_time"`
	EndTime          time.Time   `json:"end_time"`
	Description      string      `json:"description"`
	MainPOC          string      `json:"main_poc"`
	RecordAttendance bool        `json:"record_attendance"`
}

type EventCordinator struct {
	gorm.Model
	JobPerformaEventID uint             `gorm:"index" json:"job_performa_event_id"`
	JobPerformaEvent   JobPerformaEvent `gorm:"foreignkey:JobPerformaEventID" json:"-"`
	CordinatorID       string           `json:"cordinator_id"`
	CordinatorName     string           `json:"cordinator_name"`
}

type EventStudent struct {
	gorm.Model
	JobPerformaEventID        uint             `gorm:"index" json:"job_performa_event_id"`
	JobPerformaEvent          JobPerformaEvent `gorm:"foreignkey:JobPerformaEventID" json:"-"`
	StudentRecruitmentCycleID uint             `gorm:"index" json:"student_recruitment_cycle_id"`
	Present                   bool             `json:"present"`
}
