package company

import (
	"database/sql"
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
type ProformaEvent struct {
	gorm.Model
	ProformaID       uint     `json:"proforma_id" gorm:"index;->;<-:create"`
	Proforma         Proforma `json:"-" gorm:"foreignkey:ProformaID"`
	CalID            string   `json:"-"`
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
type EventStudent struct {
	gorm.Model
	ProformaEventID           uint          `json:"proforma_event_id" gorm:"index;->;<-:create"`
	ProformaEvent             ProformaEvent `json:"-" gorm:"foreignkey:ProformaEventID"`
	CompanyRecruitmentCycleID uint          `json:"company_recruitment_cycle_id" gorm:"index;->;<-:create"`
	StudentRecruitmentCycleID uint          `json:"student_recruitment_cycle_id" gorm:"index;->;<-:create"`
	Present                   bool          `json:"present" gorm:"default:true"`
}
