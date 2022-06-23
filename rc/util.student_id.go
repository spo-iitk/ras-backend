package rc

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/ras-backend/middleware"
)

func getStudentRecruitmentCycleID(ctx *gin.Context, rid uint) (uint, bool, error) {
	var student StudentRecruitmentCycle

	email := middleware.GetUserID(ctx)

	err := fetchStudent(ctx, email, rid, &student)
	if err != nil {
		return 0, false, err
	}

	return student.ID, student.IsVerified, err
}
