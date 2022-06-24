package application

import (
	"database/sql"

	"gorm.io/gorm"
)

type Proforma struct {
	gorm.Model
	Eligibility               string        `gorm:"index" json:"eligibility"`
	CPI                       float64       `json:"cpi" gorm:"default:0"`
	CompanyID                 uint          `gorm:"index" json:"company_id"`
	CompanyRecruitmentCycleID uint          `gorm:"index" json:"company_recruitment_cycle_id"`
	RecruitmentCycleID        uint          `gorm:"index" json:"recruitment_cycle_id"`
	IsApproved                sql.NullBool  `json:"is_approved" gorm:"default:false"`
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
	SHORTANSWER ApplicationQuestionsType = "Short Answer"
	BOOLEAN     ApplicationQuestionsType = "Boolean"
)

type JobApplicationQuestion struct {
	gorm.Model
	ProformaID uint                     `gorm:"index" json:"proforma_id"`
	Proforma   Proforma                 `gorm:"foreignkey:ProformaID" json:"-"`
	Type       ApplicationQuestionsType `json:"type"`
	Question   string                   `json:"question"`
	Options    string                   `json:"options"` //csv
}

type JobApplicationQuestionsAnswer struct {
	gorm.Model
	JobApplicationQuestionID  uint                   `gorm:"index" json:"job_application_question_id"`
	JobApplicationQuestion    JobApplicationQuestion `gorm:"foreignkey:JobApplicationQuestionID" json:"-"`
	StudentRecruitmentCycleID uint                   `gorm:"index" json:"student_recruitment_cycle_id"`
	Answer                    string                 `json:"answer"`
}

type ProformaEvent struct {
	gorm.Model
	ProformaID       uint     `gorm:"index" json:"proforma_id"`
	Proforma         Proforma `gorm:"foreignkey:ProformaID" json:"-"`
	Name             string   `json:"name"`
	Duration         string   `json:"duration"`
	Venue            string   `json:"venue"`
	StartTime        int64    `json:"start_time"`
	EndTime          int64    `json:"end_time"`
	Description      string   `json:"description"`
	MainPOC          string   `json:"main_poc"`
	Sequence         int      `json:"sequence"`
	RecordAttendance bool     `json:"record_attendance" gorm:"default:false"`
}

type EventCoordinator struct {
	gorm.Model
	ProformaEventID uint          `gorm:"index" json:"proforma_event_id"`
	ProformaEvent   ProformaEvent `gorm:"foreignkey:ProformaEventID" json:"-"`
	CordinatorID    string        `json:"cordinator_id"`
	CordinatorName  string        `json:"cordinator_name"`
}

type EventStudent struct {
	gorm.Model
	CompanyRecruitmentCycleID uint          `gorm:"index" json:"company_recruitment_cycle_id"`
	ProformaEventID           uint          `gorm:"index" json:"proforma_event_id"`
	ProformaEvent             ProformaEvent `gorm:"foreignkey:ProformaEventID" json:"-"`
	StudentRecruitmentCycleID uint          `gorm:"index" json:"student_recruitment_cycle_id"`
	Present                   bool          `gorm:"default:true" json:"present"`
}
