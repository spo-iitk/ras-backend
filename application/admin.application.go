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
	Name        string `json:"name"`
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
	StatusName                   string  `json:"status_name"`
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

	var allStudentsRC []rc.StudentRecruitmentCycle
	rc.FetchStudentBySRID(ctx, srids, &allStudentsRC)

	var allStudentsRCMap = make(map[uint]*rc.StudentRecruitmentCycle)
	for i := range allStudentsRC {
		allStudentsRCMap[allStudentsRC[i].ID] = &allStudentsRC[i]
	}

	var sid []uint
	for _, student := range allStudentsRC {
		sid = append(sid, student.StudentID)
	}

	var allStudents []student.Student
	student.FetchStudentsByID(ctx, sid, &allStudents)

	var allStudentsMap = make(map[uint]*student.Student)
	for i := range allStudents {
		allStudentsMap[allStudents[i].ID] = &allStudents[i]
	}

	var validApplicants []studentAdminsideResponse
	for _, student := range applied {
		if allStudentsRCMap[student.StudentRCID].IsFrozen {
			continue
		}

		applicant_details := studentAdminsideResponse{}
		applicant_details.ID = student.StudentRCID
		applicant_details.Resume = student.ResumeLink
		applicant_details.StatusName = student.Name

		sid := allStudentsRCMap[student.StudentRCID].StudentID

		applicant_details.CPI = allStudentsRCMap[student.StudentRCID].CPI

		applicant_details.Name = allStudentsMap[sid].Name
		applicant_details.Email = allStudentsMap[sid].IITKEmail
		applicant_details.ProgramDepartmentID = allStudentsMap[sid].ProgramDepartmentID
		applicant_details.SecondaryProgramDepartmentID = allStudentsMap[sid].SecondaryProgramDepartmentID
		applicant_details.CurrentCPI = allStudentsMap[sid].CurrentCPI
		applicant_details.UGCPI = allStudentsMap[sid].UGCPI
		applicant_details.TenthBoard = allStudentsMap[sid].TenthBoard
		applicant_details.TenthYear = allStudentsMap[sid].TenthYear
		applicant_details.TenthMarks = allStudentsMap[sid].TenthMarks
		applicant_details.TwelfthBoard = allStudentsMap[sid].TwelfthBoard
		applicant_details.TwelfthYear = allStudentsMap[sid].TwelfthYear
		applicant_details.TwelfthMarks = allStudentsMap[sid].TwelfthMarks
		applicant_details.EntranceExam = allStudentsMap[sid].EntranceExam
		applicant_details.EntranceExamRank = allStudentsMap[sid].EntranceExamRank
		applicant_details.Category = allStudentsMap[sid].Category
		applicant_details.CategoryRank = allStudentsMap[sid].CategoryRank
		applicant_details.CurrentAddress = allStudentsMap[sid].CurrentAddress
		applicant_details.PermanentAddress = allStudentsMap[sid].PermanentAddress
		applicant_details.FriendName = allStudentsMap[sid].FriendName
		applicant_details.FriendPhone = allStudentsMap[sid].FriendPhone

		validApplicants = append(validApplicants, applicant_details)
	}

	ctx.JSON(http.StatusOK, validApplicants)
}
