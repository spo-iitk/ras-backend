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
	Profile                   string `json:"profile"`
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
	Profile                      string `json:"profile"`
	Type                         string `json:"type"`
}

type companyRecruitResponce struct {
	CompanyRecruitmentCycleID uint `json:"company_recruitment_cycle_id"`
	Count                     int  `json:"count"`
}

func fetchCompanyRecruitCountHandler(c *gin.Context) {
	var cids []uint

	if err := c.ShouldBindJSON(&cids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	recruitCounts, err := fetchCompanyRecruitCount(c, cids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ppoCounts, err := fetchCompanyPPOCount(c, cids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recruitCounts": recruitCounts, "ppoCount": ppoCounts})
}

func getCompanyRecruitStatsHandler(ctx *gin.Context) {
	cid, err := util.ParseUint(ctx.Param("cid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var stats []statsResponse
	err = getCompanyRecruitmentStats(ctx, cid, &stats)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var srids []uint
	for _, stat := range stats {
		srids = append(srids, stat.StudentRecruitmentCycleID)
	}

	var students []rc.StudentRecruitmentCycle
	err = rc.FetchStudentBySRID(ctx, srids, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentsMap = make(map[uint]*rc.StudentRecruitmentCycle)
	for i := range students {
		studentsMap[students[i].ID] = &students[i]
	}

	var studentResponse []statsRecruitmentResponse
	for _, stat := range stats {
		student := studentsMap[stat.StudentRecruitmentCycleID]
		res := statsRecruitmentResponse{
			ID:                           student.ID,
			Name:                         student.Name,
			Email:                        student.Email,
			RollNo:                       student.RollNo,
			ProgramDepartmentID:          student.ProgramDepartmentID,
			SecondaryProgramDepartmentID: student.SecondaryProgramDepartmentID,
			CompanyName:                  stat.CompanyName,
			Profile:                      stat.Profile,
			Type:                         stat.Type,
		}
		studentResponse = append(studentResponse, res)
	}

	ctx.JSON(http.StatusOK, gin.H{"student": studentResponse})

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

	var branchStats []rc.StatsBranchResponse
	err = rc.FetchRegisteredStudentCountByBranch(ctx, rid, &branchStats)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var secondaryBranchStats []rc.StatsBranchResponse
	err = rc.FetchRegisteredStudentCountBySecondaryBranch(ctx, rid, &secondaryBranchStats)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, secondary := range secondaryBranchStats {
		found := false
		for i, primary := range branchStats {
			if primary.ProgramDepartmentID == secondary.ProgramDepartmentID {
				branchStats[i].Total += secondary.Total
				found = true
				break
			}
		}
		if !found {
			branchStats = append(branchStats, secondary)
		}
	}

	var srids []uint
	for _, stat := range stats {
		srids = append(srids, stat.StudentRecruitmentCycleID)
	}

	var students []rc.StudentRecruitmentCycle
	err = rc.FetchStudentBySRID(ctx, srids, &students)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentsMap = make(map[uint]*rc.StudentRecruitmentCycle)
	for i := range students {
		studentsMap[students[i].ID] = &students[i]
	}

	var branchMap = make(map[uint]*rc.StatsBranchResponse)
	for i := range branchStats {
		branchMap[branchStats[i].ProgramDepartmentID] = &branchStats[i]
	}

	var studentResponse []statsRecruitmentResponse
	for _, stat := range stats {
		student := studentsMap[stat.StudentRecruitmentCycleID]
		res := statsRecruitmentResponse{
			ID:                           student.ID,
			Name:                         student.Name,
			Email:                        student.Email,
			RollNo:                       student.RollNo,
			ProgramDepartmentID:          student.ProgramDepartmentID,
			SecondaryProgramDepartmentID: student.SecondaryProgramDepartmentID,
			CompanyName:                  stat.CompanyName,
			Profile:                      stat.Profile,
			Type:                         stat.Type,
		}
		studentResponse = append(studentResponse, res)

		if res.Type == string(Recruited) {
			branchMap[res.ProgramDepartmentID].Recruited++
			if res.SecondaryProgramDepartmentID != 0 {
				branchMap[res.SecondaryProgramDepartmentID].Recruited++
			}
		}

		if res.Type == string(PIOPPOACCEPTED) {
			branchMap[res.ProgramDepartmentID].PreOffer++
			if res.SecondaryProgramDepartmentID != 0 {
				branchMap[res.SecondaryProgramDepartmentID].Recruited++
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"student": studentResponse, "branch": branchStats})
}
