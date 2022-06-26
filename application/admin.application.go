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

	var allStudents []rc.StudentRecruitmentCycle
	rc.FetchStudentBySRID(ctx, srids, &allStudents)

	var allStudentsMap = make(map[uint]*rc.StudentRecruitmentCycle)
	for i := range allStudents {
		allStudentsMap[allStudents[i].ID] = &allStudents[i]
	}

	var sid []uint
	for _, student := range allStudents {
		sid = append(sid, student.StudentID)
	}

	var allStudentDetails []student.Student
	student.FetchStudentsByID(ctx, sid, &allStudentDetails)

	var allStudentDetailsMap = make(map[uint]*student.Student)
	for i := range allStudentDetails {
		allStudentDetailsMap[allStudentDetails[i].ID] = &allStudentDetails[i]
	}

	var validApplicants []studentAdminsideResponse
	for _, student := range applied {
		if allStudentsMap[student.StudentRCID].IsFrozen {
			continue
		}

		applicant_details := studentAdminsideResponse{}
		applicant_details.ID = student.StudentRCID
		applicant_details.Resume = student.ResumeLink
		applicant_details.StatusName = student.Name

		sid := allStudentsMap[student.StudentRCID].StudentID

		applicant_details.CPI = allStudentsMap[student.StudentRCID].CPI

		applicant_details.Name = allStudentDetailsMap[sid].Name
		applicant_details.Email = allStudentDetailsMap[sid].IITKEmail
		applicant_details.ProgramDepartmentID = allStudentDetailsMap[sid].ProgramDepartmentID
		applicant_details.SecondaryProgramDepartmentID = allStudentDetailsMap[sid].SecondaryProgramDepartmentID
		applicant_details.CurrentCPI = allStudentDetailsMap[sid].CurrentCPI
		applicant_details.UGCPI = allStudentDetailsMap[sid].UGCPI
		applicant_details.TenthBoard = allStudentDetailsMap[sid].TenthBoard
		applicant_details.TenthYear = allStudentDetailsMap[sid].TenthYear
		applicant_details.TenthMarks = allStudentDetailsMap[sid].TenthMarks
		applicant_details.TwelfthBoard = allStudentDetailsMap[sid].TwelfthBoard
		applicant_details.TwelfthYear = allStudentDetailsMap[sid].TwelfthYear
		applicant_details.TwelfthMarks = allStudentDetailsMap[sid].TwelfthMarks
		applicant_details.EntranceExam = allStudentDetailsMap[sid].EntranceExam
		applicant_details.EntranceExamRank = allStudentDetailsMap[sid].EntranceExamRank
		applicant_details.Category = allStudentDetailsMap[sid].Category
		applicant_details.CategoryRank = allStudentDetailsMap[sid].CategoryRank
		applicant_details.CurrentAddress = allStudentDetailsMap[sid].CurrentAddress
		applicant_details.PermanentAddress = allStudentDetailsMap[sid].PermanentAddress
		applicant_details.FriendName = allStudentDetailsMap[sid].FriendName
		applicant_details.FriendPhone = allStudentDetailsMap[sid].FriendPhone

		validApplicants = append(validApplicants, applicant_details)
	}

	ctx.JSON(http.StatusOK, validApplicants)
}
