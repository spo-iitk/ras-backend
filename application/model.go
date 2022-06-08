package application

import (
	"database/sql"

	"gorm.io/gorm"
)

type JobProforma struct {
	gorm.Model
	CompanyID                 uint          `gorm:"index" json:"company_id"`
	CompanyRecruitmentCycleID uint          `gorm:"index" json:"company_recruitment_cycle_id"`
	RecruitmentCycleID        uint          `gorm:"index" json:"recruitment_cycle_id"`
	IsApproved                bool          `json:"is_approved"`
	ActionTakenBy             string        `json:"action_taken_by"`
	SetDeadline               sql.NullInt64 `json:"set_deadline"` // NULL implies unpublished
	HideDetails               bool          `gorm:"default:false" json:"hide_details"`
	ActiveHRID                string        `json:"active_hr_id"`
	NatureOfBusiness          string        `json:"nature_of_business"`
	TentativeJobLocation      string        `json:"tentative_job_location"`
	JobDescription            string        `json:"job_description"`
	CostToCompany             string        `json:"cost_to_company"`
	PackageDetails            string        `json:"package_details"`
	BondDetails               string        `json:"bond_details"`
	MedicalRequirements       string        `json:"medical_requirements"`
	AdditionalEligibility     string        `json:"additional_eligibility"`
	MessageForCordinator      string        `json:"message_for_cordinator"`
}

type ApplicationQuestionsType string

const (
	MCQ         ApplicationQuestionsType = "MCQ"
	SHORTANSWER ApplicationQuestionsType = "ShortAnswer"
	BOOLEAN     ApplicationQuestionsType = "Boolean"
)

type JobApplicationQuestion struct {
	gorm.Model
	JobPerformaID uint                     `gorm:"index" json:"job_performa_id"`
	JobPerforma   JobProforma              `gorm:"foreignkey:JobPerformaID" json:"-"`
	Type          ApplicationQuestionsType `json:"type"`
	Question      string                   `json:"question"`
	Options       string                   `json:"options"` //csv
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
	StartTime        int64       `json:"start_time"`
	EndTime          int64       `json:"end_time"`
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
