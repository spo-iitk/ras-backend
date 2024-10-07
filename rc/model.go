package rc

import (
	"database/sql"

	"gorm.io/gorm"
)

type RecruitmentCycleType string

const (
	PLACEMENT  RecruitmentCycleType = "Placement"
	INTERNSHIP RecruitmentCycleType = "Internship"
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
	PostalAddress             string       `json:"postal_address"`
	EstablishmentDate         string       `json:"establishment_date"`
	TotalEmployees            string       `json:"total_employees"`
	SocialMedia               string       `json:"social_media"`
	Website                   string       `json:"website"`
	Turnover                  string       `json:"turnover"`
	TypeOfOrg                 string       `json:"type_of_org"`
	HeadOffice                string       `json:"head_office"`
	MinHires                  string       `json:"min_hires"`
	TotalHires                string       `json:"total_hires"`
	SkillSet                  string       `json:"skill_set"`
	PWD                       string       `json:"pwd"`
	CPICriteria               string       `json:"cpi_criteria"`
	BacklogEligibility        string       `json:"backlog_eligibility"`
	CTCINR                    string       `json:"ctc_inr"`
	CTCFR                     string       `json:"ctc_fr"`
	Gross                     string       `json:"gross"`
	TakeHome                  string       `json:"take_home"`
	Base                      string       `json:"base"`
	JoiningBonus              string       `json:"joining_bonus"`
	RelocationBonus           string       `json:"relocation_bonus"`
	FirstCTC                  string       `json:"first_ctc"`
	MedicalAllowance          string       `json:"medical_allowance"`
	RetentionBonus            string       `json:"retention_bonus"`
	Deductions                string       `json:"deductions"`
	Perks                     string       `json:"perks"`
	Accommodation             string       `json:"accommodation"`
	PPOConfirmingDate         string       `json:"ppo_confirming_date"`
	InternshipPeriod          string       `json:"internship_period"`
}

type RecruitmentCycle struct {
	gorm.Model
	IsActive            bool                 `json:"is_active" gorm:"default:true"`
	AcademicYear        string               `json:"academic_year"`
	Type                RecruitmentCycleType `json:"type"`
	StartDate           int64                `json:"start_date"`
	Phase               string               `json:"phase"`
	ApplicationCountCap uint                 `json:"application_count_cap"`
}

type RecruitmentCycleQuestionsType string

const (
	MCQ         RecruitmentCycleQuestionsType = "MCQ"
	SHORTANSWER RecruitmentCycleQuestionsType = "Short Answer"
	BOOLEAN     RecruitmentCycleQuestionsType = "Boolean"
)

type RecruitmentCycleQuestion struct {
	gorm.Model
	Type               RecruitmentCycleQuestionsType `json:"type"`
	Question           string                        `json:"question"`
	RecruitmentCycleID uint                          `gorm:"index;->;<-:create" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle              `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Mandatory          bool                          `json:"mandatory" gorm:"default:false"`
	Options            string                        `json:"options"` //csv
}

type RecruitmentCycleQuestionsAnswer struct {
	gorm.Model
	RecruitmentCycleQuestionID uint                     `gorm:"index;->;<-:create" json:"recruitment_cycle_question_id"`
	RecruitmentCycleQuestion   RecruitmentCycleQuestion `gorm:"foreignkey:RecruitmentCycleQuestionID" json:"-"`
	StudentRecruitmentCycleID  uint                     `gorm:"index;->;<-:create" json:"student_recruitment_cycle_id"`
	StudentRecruitmentCycle    StudentRecruitmentCycle  `gorm:"foreignkey:StudentRecruitmentCycleID" json:"-"`
	Answer                     string                   `json:"answer"`
}

type CompanyRecruitmentCycle struct {
	gorm.Model
	CompanyID          uint             `gorm:"index" json:"company_id"`
	CompanyName        string           `json:"company_name"`
	RecruitmentCycleID uint             `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	HR1                string           `json:"hr1"`
	HR2                string           `json:"hr2"`
	HR3                string           `json:"hr3"`
	Comments           string           `json:"comments"`
}

type Notice struct {
	gorm.Model
	RecruitmentCycleID uint             `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle   RecruitmentCycle `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Title              string           `json:"title" binding:"required"`
	Description        string           `json:"description" binding:"required"`
	Tags               string           `json:"tags" binding:"required"`
	Attachment         string           `json:"attachment"`
	CreatedBy          string           `json:"created_by"`
	LastReminderAt     int64            `json:"last_reminder_at" gorm:"default:0"`
	Deadline           uint             `json:"deadline" gorm:"default:0"`
}

type StudentRecruitmentCycleType string

const (
	PIOPPO       StudentRecruitmentCycleType = "PIO-PPO"
	RECRUITED    StudentRecruitmentCycleType = "Recruited"
	AVAILABLE    StudentRecruitmentCycleType = "Available"
	DEREGISTERED StudentRecruitmentCycleType = "Deregistered"
)

type StudentRecruitmentCycle struct {
	gorm.Model
	StudentID                    uint                        `gorm:"index:stu_rec_cycle,unique" json:"student_id"`
	RecruitmentCycleID           uint                        `gorm:"index:stu_rec_cycle,unique" json:"recruitment_cycle_id"`
	RecruitmentCycle             RecruitmentCycle            `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	ProgramDepartmentID          uint                        `gorm:"index" json:"program_department_id"`
	SecondaryProgramDepartmentID uint                        `gorm:"index" json:"secondary_program_department_id"`
	CPI                          float64                     `json:"cpi"`
	Email                        string                      `gorm:"index" json:"email"`
	RollNo                       string                      `gorm:"index" json:"roll_no"`
	Name                         string                      `json:"name"`
	Type                         StudentRecruitmentCycleType `json:"type" gorm:"default:Available"`
	IsFrozen                     bool                        `json:"is_frozen" gorm:"default:false"`
	IsVerified                   bool                        `json:"is_verified" gorm:"default:false"`
	Comment                      string                      `json:"comment"`
}
type ResumeType string

const (
	SINGLE ResumeType = "Single"
	MASTER ResumeType = "Master"
)

type StudentRecruitmentCycleResume struct {
	gorm.Model
	StudentRecruitmentCycleID uint                    `gorm:"index" json:"student_recruitment_cycle_id"`
	StudentRecruitmentCycle   StudentRecruitmentCycle `gorm:"foreignkey:StudentRecruitmentCycleID" json:"-"`
	RecruitmentCycleID        uint                    `gorm:"index" json:"recruitment_cycle_id"`
	RecruitmentCycle          RecruitmentCycle        `gorm:"foreignkey:RecruitmentCycleID" json:"-"`
	Resume                    string                  `json:"resume"`
	Verified                  sql.NullBool            `json:"verified" gorm:"default:NULL"`
	ActionTakenBy             string                  `json:"action_taken_by"`
	ResumeType                ResumeType              `json:"resume_type"` // New field
}

type CompanyHistory struct {
	ID                 uint   `json:"id" gorm:"column:id"`
	RecruitmentCycleID uint   `json:"recruitment_cycle_id" gorm:"column:recruitment_cycle_id"`
	Type               string `json:"type" gorm:"column:type"`
	Phase              string `json:"phase" gorm:"column:phase"`
	Comments           string `json:"comments" gorm:"column:comments"`
}

type CompanyHistoryResponse struct {
	ID                 uint   `json:"id" gorm:"column:id"`
	RecruitmentCycleID uint   `json:"recruitmentCycleID" gorm:"column:recruitment_cycle_id"`
	Comments           string `json:"comments" gorm:"column:comments"`
}
