package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/rc"
	"github.com/spo-iitk/ras-backend/util"
)

type ApplicantsByRole struct {
	StudentID  uint   `json:"student_id"`
	ResumeLink string `json:"resume_link"`
	Status     string `json:"status"`
}

type studentResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Resume string `json:"resume"`
	Status string `json:"status"`
}

func getStudentsByRole(ctx *gin.Context) {
	pid, err := util.ParseUint(ctx.Param("pid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rid := ctx.Param("rid")

	var allStudents []rc.StudentRecruitmentCycle
	rc.FetchAllStudents(ctx, rid, &allStudents)

	var applied []ApplicantsByRole
	fetchApplicantDetails(ctx, pid, &applied)

	var validApplicants []studentResponse

	for _, student := range applied {
		for _, s := range allStudents {
			if s.ID == student.StudentID && !s.IsFrozen {
				validApplicants = append(validApplicants, studentResponse{
					ID:     s.ID,
					Name:   s.Name,
					Email:  s.Email,
					Resume: student.ResumeLink,
					Status: student.Status,
				})
			}
		}
	}

	ctx.JSON(http.StatusOK, validApplicants)
}
