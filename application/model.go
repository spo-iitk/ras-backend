package application

import (
	"database/sql"

	"gorm.io/gorm"
)

type Proforma struct {
	gorm.Model
	CompanyRecruitmentCycleID uint         `json:"company_recruitment_cycle_id" gorm:"index;->;<-:create"`
	RecruitmentCycleID        uint         `json:"recruitment_cycle_id" gorm:"index;->;<-:create"`
	CompanyID                 uint         `json:"company_id" gorm:"index;->;<-:create"`
	CompanyName               string       `json:"company_name"`
	ActionTakenBy             string       `json:"action_taken_by"`
	IsApproved                sql.NullBool `json:"is_approved" gorm:"index;default:NULL"`
	Deadline                  uint         `json:"deadline" gorm:"default:0"` // 0 implies unpublished
	Eligibility               string       `json:"eligibility"`
	CPICutoff                 float64      `json:"cpi_cutoff" gorm:"default:0"`
	HideDetails               bool         `json:"hide_details" gorm:"default:true"`
	ActiveHR                  string       `json:"active_hr"`
	Role                      string       `json:"role"`
	Profile                   string       `json:"profile"`
	TentativeJobLocation      string       `json:"tentative_job_location"`
	JobDescription            string       `json:"job_description"`
	CostToCompany             string       `json:"cost_to_company"`
	PackageDetails            string       `json:"package_details"`
	BondDetails               string       `json:"bond_details"`
	MedicalRequirements       string       `json:"medical_requirements"`
	AdditionalEligibility     string       `json:"additional_eligibility"`
	MessageForCordinator      string       `json:"message_for_cordinator"`
}

type ApplicationQuestionType string

const (
	MCQ         ApplicationQuestionType = "MCQ"
	SHORTANSWER ApplicationQuestionType = "Short Answer"
	BOOLEAN     ApplicationQuestionType = "Boolean"
)

type ApplicationQuestion struct {
	gorm.Model
	ProformaID uint                    `json:"proforma_id" gorm:"index;->;<-:create"`
	Proforma   Proforma                `json:"-" gorm:"foreignkey:ProformaID"`
	Type       ApplicationQuestionType `json:"type"`
	Question   string                  `json:"question"`
	Options    string                  `json:"options"` //csv
}

type ApplicationQuestionAnswer struct {
	gorm.Model
	ApplicationQuestionID     uint                `json:"application_question_id" gorm:"index;->;<-:create"`
	ApplicationQuestion       ApplicationQuestion `json:"-" gorm:"foreignkey:ApplicationQuestionID"`
	StudentRecruitmentCycleID uint                `json:"student_recruitment_cycle_id" gorm:"index;->;<-:create"`
	Answer                    string              `json:"answer"`
}

type ProformaEvent struct {
	gorm.Model
	ProformaID       uint     `json:"proforma_id" gorm:"index;->;<-:create"`
	Proforma         Proforma `json:"-" gorm:"foreignkey:ProformaID"`
	CalID            string   `json:"cal_id"`
	Name             string   `json:"name"`
	Duration         string   `json:"duration"`
	Venue            string   `json:"venue"`
	StartTime        uint     `json:"start_time"`
	EndTime          uint     `json:"end_time"`
	Description      string   `json:"description"`
	MainPOC          string   `json:"main_poc"`
	Sequence         int      `json:"sequence"`
	RecordAttendance bool     `json:"record_attendance" gorm:"default:false"`
}

type EventCoordinator struct {
	gorm.Model
	ProformaEventID uint          `json:"proforma_event_id" gorm:"index;->;<-:create"`
	ProformaEvent   ProformaEvent `json:"-" gorm:"foreignkey:ProformaEventID"`
	CordinatorID    string        `json:"cordinator_id"`
	CordinatorName  string        `json:"cordinator_name"`
}

type EventStudent struct {
	gorm.Model
	ProformaEventID           uint          `json:"proforma_event_id" gorm:"index;->;<-:create"`
	ProformaEvent             ProformaEvent `json:"-" gorm:"foreignkey:ProformaEventID"`
	CompanyRecruitmentCycleID uint          `json:"company_recruitment_cycle_id" gorm:"index;->;<-:create"`
	StudentRecruitmentCycleID uint          `json:"student_recruitment_cycle_id" gorm:"index;->;<-:create"`
	Present                   bool          `json:"present" gorm:"default:true"`
}

type ApplicationResume struct {
	gorm.Model
	StudentRecruitmentCycleID uint     `json:"student_recruitment_cycle_id" gorm:"index;->;<-:create"`
	ProformaID                uint     `json:"proforma_id" gorm:"index;->;<-:create"`
	Profoma                   Proforma `json:"-" gorm:"foreignkey:ProformaID"`
	ResumeID                  uint     `json:"resume_id"`
	Resume                    string   `json:"resume"`
}
