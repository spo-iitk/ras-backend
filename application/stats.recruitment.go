package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

type statsResponse struct {
	StudentRecruitmentCycleID uint   `json:"student_recruitment_cycle_id"`
	CompanyName               string `json:"company_name"`
	Role                      string `json:"role"`
	Type                      string `json:"type"`
}

type statsRecruitmentResponse struct {
	ID                           uint   `json:"id"`
	Name                         string `json:"name"`
	Email                        string `json:"email"`
	RollNo                       string `json:"roll_no"`
	ProgramDepartmentID          uint   `json:"program_department_id"`
	SecondaryProgramDepartmentID uint   `json:"secondary_program_department_id"`
	CompanyName                  string `json:"company_name"`
	Role                         string `json:"role"`
	Type                         string `json:"type"`
}

func getStatsHandler(ctx *gin.Context) {
	rid, err := util.ParseUint(ctx.Param("rid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var stats []statsResponse
	err = getRecruitmentStats(ctx, rid, &stats)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var srids []uint
	for _, stat := range stats {
		srids = append(srids, stat.StudentRecruitmentCycleID)
	}

	var students []rc.StudentRecruitmentCycle
	rc.FetchStudentBySRID(ctx, srids, &students)

	var studentsMap = make(map[uint]*rc.StudentRecruitmentCycle)
	for i := range students {
		studentsMap[students[i].ID] = &students[i]
	}

	var response []statsRecruitmentResponse
	for _, stat := range stats {
		student := studentsMap[stat.StudentRecruitmentCycleID]
		res := statsRecruitmentResponse{
			ID:                           student.ID,
			Name:                         student.Name,
			Email:                        student.Email,
			RollNo:                       student.Email,
			ProgramDepartmentID:          student.ProgramDepartmentID,
			SecondaryProgramDepartmentID: student.SecondaryProgramDepartmentID,
			CompanyName:                  stat.CompanyName,
			Role:                         stat.Role,
			Type:                         stat.Type,
		}
		response = append(response, res)
	}

	ctx.JSON(http.StatusOK, response)
}
