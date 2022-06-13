package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentRecruitmentCycleID(ctx *gin.Context, rid string) (uint, error) {
	var student StudentRecruitmentCycle

	email := middleware.GetUserID(ctx)

	err := fetchStudent(ctx, email, rid, &student)
	if err != nil {
		return 0, err
	}

	return student.ID, err
}
