package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/student"
	"github.com/spo-iitk/ras-backend/util"
)

type ApplicantsByRole struct {
	StudentRCID uint   `json:"student_rc_id"`
	ResumeLink  string `json:"resume_link"`
	ProformaID  uint   `json:"proforma_id"`
	Status      string `json:"status"`
}

type studentAdminsideResponse struct {
	ID                           uint    `json:"id"`
	Name                         string  `json:"name"`
	Email                        string  `json:"email"`
	CPI                          float64 `json:"cpi"`
	ProgramDepartmentID          uint    `json:"program_department_id"`
	SecondaryProgramDepartmentID uint    `json:"secondary_program_department_id"`
	CurrentCPI                   float64 `json:"current_cpi"`
	UGCPI                        float64 `json:"ug_cpi"`
	TenthBoard                   string  `json:"tenth_board"`
	TenthYear                    uint    `json:"tenth_year"`
	TenthMarks                   float64 `json:"tenth_marks"`
	TwelfthBoard                 string  `json:"twelfth_board"`
	TwelfthYear                  uint    `json:"twelfth_year"`
	TwelfthMarks                 float64 `json:"twelfth_marks"`
	EntranceExam                 string  `json:"entrance_exam"`
	EntranceExamRank             uint    `json:"entrance_exam_rank"`
	Category                     string  `json:"category"`
	CategoryRank                 uint    `json:"category_rank"`
	CurrentAddress               string  `json:"current_address"`
	PermanentAddress             string  `json:"permanent_address"`
	FriendName                   string  `json:"friend_name"`
	FriendPhone                  string  `json:"friend_phone"`
	Resume                       string  `json:"resume"`
	Status                       string  `json:"status"`
}

func getStudentsByRole(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var applied []ApplicantsByRole
	fetchApplicantDetails(ctx, pid, &applied)

	var srids []uint
	for _, applicant := range applied {
		srids = append(srids, applicant.StudentRCID)
	}

	var allStudents []rc.StudentRecruitmentCycle
	rc.FetchStudentBySRID(ctx, srids, &allStudents)

	var sid []uint
	for _, student := range allStudents {
		sid = append(sid, student.StudentID)
	}

	var allStudentDetails []student.Student
	student.FetchStudentsByID(ctx, sid, &allStudentDetails)

	var validApplicants []studentAdminsideResponse
	// for _, student := range applied {
	// }

	ctx.JSON(http.StatusOK, validApplicants)
}
